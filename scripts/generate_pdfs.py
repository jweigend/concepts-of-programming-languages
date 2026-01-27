#!/usr/bin/env python3
"""
Generate PDFs from Go present slides.
Requires: selenium, webdriver-manager, Chrome/Chromium

Usage:
    python3 scripts/generate_pdfs.py [--port 3999]

Prerequisites:
    1. Start the present server: make slideshow
    2. Run this script in another terminal
"""

import os
import sys
import time
import json
import base64
import glob
import argparse
from pathlib import Path

try:
    from selenium import webdriver
    from selenium.webdriver.chrome.service import Service
    from selenium.webdriver.chrome.options import Options
    from selenium.webdriver.common.by import By
    from selenium.webdriver.support.ui import WebDriverWait
    from selenium.webdriver.support import expected_conditions as EC
    from webdriver_manager.chrome import ChromeDriverManager
except ImportError:
    print("Error: selenium and webdriver-manager are required.")
    print("Install with: pip3 install selenium webdriver-manager --break-system-packages")
    sys.exit(1)


def setup_driver():
    """Setup Chrome driver with PDF printing capabilities."""
    options = Options()
    options.add_argument("--headless=new")
    options.add_argument("--disable-gpu")
    options.add_argument("--no-sandbox")
    options.add_argument("--disable-dev-shm-usage")
    options.add_argument("--window-size=1920,1080")
    
    # Enable PDF printing
    options.add_experimental_option("prefs", {
        "printing.print_preview_sticky_settings.appState": json.dumps({
            "recentDestinations": [{"id": "Save as PDF", "origin": "local", "account": ""}],
            "selectedDestinationId": "Save as PDF",
            "version": 2,
            "isLandscapeEnabled": True,
            "pageSize": "A4",
            "marginsType": 0,
            "scalingType": 0,
            "scaling": "100",
        })
    })
    
    service = Service(ChromeDriverManager().install())
    driver = webdriver.Chrome(service=service, options=options)
    return driver


def save_as_pdf(driver, output_path, scale=0.8):
    """Save current page as PDF using Chrome DevTools Protocol."""
    # Use CDP to print to PDF with reduced scale for smaller file size
    result = driver.execute_cdp_cmd("Page.printToPDF", {
        "landscape": True,
        "printBackground": True,
        "preferCSSPageSize": False,
        "paperWidth": 11,  # inches (Letter landscape width)
        "paperHeight": 8.5,  # inches (Letter landscape height)
        "marginTop": 0,
        "marginBottom": 0,
        "marginLeft": 0,
        "marginRight": 0,
        "scale": scale,  # Reduce scale for smaller file size
    })
    
    # Decode and save
    pdf_data = base64.b64decode(result["data"])
    with open(output_path, "wb") as f:
        f.write(pdf_data)
    
    return len(pdf_data)


def compress_pdf(input_path, output_path):
    """Compress PDF using Ghostscript if available."""
    import subprocess
    try:
        # Use Ghostscript for compression
        subprocess.run([
            "gs", "-sDEVICE=pdfwrite", "-dCompatibilityLevel=1.4",
            "-dPDFSETTINGS=/ebook",  # Good quality, smaller size
            "-dNOPAUSE", "-dQUIET", "-dBATCH",
            f"-sOutputFile={output_path}",
            str(input_path)
        ], check=True, capture_output=True)
        return True
    except (subprocess.CalledProcessError, FileNotFoundError):
        return False


def count_slides(driver):
    """Count total number of slides."""
    try:
        # The present tool shows slide count in the URL hash or we can navigate
        driver.execute_script("window.location.hash = '#999';")
        time.sleep(0.3)
        # Get the actual last slide number from the hash
        current_hash = driver.execute_script("return window.location.hash;")
        if current_hash:
            return int(current_hash.replace("#", ""))
    except:
        pass
    return 1


def generate_pdf_for_slide(driver, base_url, slide_file, output_dir, wait_time=3):
    """Generate PDF for a single slide file using print mode."""
    slide_url = f"{base_url}/{slide_file}?print"
    output_path = output_dir / f"{Path(slide_file).stem}.pdf"
    temp_path = output_dir / f"{Path(slide_file).stem}.tmp.pdf"
    
    print(f"  {slide_file}...", end=" ", flush=True)
    driver.get(slide_url)
    time.sleep(wait_time)
    
    # Generate PDF
    result = driver.execute_cdp_cmd("Page.printToPDF", {
        "landscape": True,
        "printBackground": True,
        "preferCSSPageSize": False,
        "paperWidth": 11,
        "paperHeight": 8.5,
        "marginTop": 0,
        "marginBottom": 0,
        "marginLeft": 0,
        "marginRight": 0,
    })
    
    pdf_data = base64.b64decode(result["data"])
    with open(temp_path, "wb") as f:
        f.write(pdf_data)
    
    # Compress with Ghostscript
    if compress_pdf(temp_path, output_path):
        os.remove(temp_path)
        size = os.path.getsize(output_path)
    else:
        os.rename(temp_path, output_path)
        size = os.path.getsize(output_path)
    
    print(f"✓ ({size // 1024} KB)")
    return True


def main():
    parser = argparse.ArgumentParser(description="Generate PDFs from Go present slides")
    parser.add_argument("--port", type=int, default=3999, help="Present server port")
    parser.add_argument("--wait", type=int, default=3, help="Wait time for page load (seconds)")
    parser.add_argument("--base-path", type=str, default="", help="Base path for slides (e.g., 'docs')")
    parser.add_argument("--compress", action="store_true", help="Compress PDFs with Ghostscript")
    args = parser.parse_args()
    
    base_url = f"http://127.0.0.1:{args.port}"
    if args.base_path:
        base_url = f"{base_url}/{args.base_path}"
    docs_dir = Path(__file__).parent.parent / "docs"
    
    # Find all slide files
    slide_files = sorted(glob.glob(str(docs_dir / "*.slide")))
    
    if not slide_files:
        print(f"Error: No .slide files found in {docs_dir}")
        sys.exit(1)
    
    print(f"Found {len(slide_files)} slide files")
    print(f"Using present server at {base_url}")
    print()
    
    # Setup driver
    print("Starting Chrome...")
    driver = setup_driver()
    
    try:
        # Test connection
        try:
            driver.get(base_url)
            time.sleep(1)
        except Exception as e:
            print(f"\nError: Cannot connect to {base_url}")
            print("Make sure the present server is running: make slideshow")
            sys.exit(1)
        
        success_count = 0
        for slide_path in slide_files:
            slide_file = os.path.basename(slide_path)
            if generate_pdf_for_slide(driver, base_url, slide_file, docs_dir, args.wait):
                success_count += 1
            print()
        
        print(f"Done! Generated {success_count}/{len(slide_files)} PDFs in {docs_dir}")
        
    finally:
        driver.quit()


if __name__ == "__main__":
    main()
