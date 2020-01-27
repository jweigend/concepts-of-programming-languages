# (Prototype) Raft Consensus Algorithm in Scala

**Protoype [Raft Consensus](https://raft.github.io/raft.pdf) Algorithm in Scala**

![](./docImg/logos.png)

Tested on ``macOs 10.15.2`` with ``openjdk64-11.0.2`` and ``sbt 1.3.3``

[![shields.io](http://img.shields.io/badge/license-Apache2-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0.txt)
![](https://github.com/maxbundscherer/prototype-scala-raft/workflows/CI%20Test/badge.svg)

Test line-coverage: 88,11% ([12-30-2019](./docImg/test-report-12-30-2019.zip))

Author: [Maximilian Bundscherer](https://bundscherer-online.de)

## Let's get started

- [sbt](https://www.scala-sbt.org/) and [openjdk64-11.0.2](https://jdk.java.net/archive/) are required to build and run project

- Run with: ``sbt run`` (see ***What happens in normal run?*** below)
- Test with: ``sbt test`` (or see ci-tests in GitHub-Actions-CI-Pipeline) (see ***What happens in test run?*** below)
- Generate test-coverage-html-report with: ``sbt jacoco``

### Used dependencies

- [akka actors](https://doc.akka.io/docs/akka/current/actors.html): Actor model implementation (Scala/Java).
- [scalactic](http://www.scalactic.org/): Test kit for Scala.
- [sbt-jacoco](https://github.com/sbt/sbt-jacoco): SBT plugin for generating coverage-reports.

### What is implemented?

- RaftNode as Finite-state machine (**FSM**) with **key-value storage**

    - ``(Uninitialized)``: Not initialized
    - ``Follower`` (Default behavior): Waiting for heartbeats from leader-node with hashCode from data. If local stored data's hashCode is not equal to leader-node data's hashCode, the node will synchronize with leader-node. If there is no heartbeat from leader-node in configured randomized interval received, the node will change to candidate-behavior. 
    - ``Candidate``: The candidate requests votes from all followers and votes for himself. If he wins the election in configured interval, he will become the leader. If not, he will become follower again. For winning the election the node requires the majority of votes.
    - ``Leader``: The leader is sending continuous heartbeats to all followers with hashCode from his stored data. The leader is the only node that is allowed to write data.
    - ``(Sleep)``: Is used for simulating leader-crashes (triggered by crashIntervalHeartbeats in normal run or by SimulateLeaderCrash in test run). In this behavior, the node does not respond to non-debug-messages. After configured downtime, the node changes to follower-behavior.
    
![](./docImg/raftFsm.png)

#### Configuration

There are two configurations:

- ``./src/main/resources/application.conf`` used for normal run
- ``./src/test/resources/application.conf`` used for test run
    
```
akka {

    # Log Level (DEBUG, INFO, WARNING, ERROR)
    loglevel = "DEBUG"

}

raftPrototype {

    # Election Timer Min (Seconds)
    electionTimerIntervalMin = 2

    # Election Timer Max (Seconds)
    electionTimerIntervalMax = 3

    # Heartbeat Timer Interval (Seconds)
    heartbeatTimerInterval = 1

    # Raft Nodes (Amount)
    nodes = 5

    # Crash Interval (auto simulate crash after some heartbeats in LEADER behavior)
    crashIntervalHeartbeats = 10

    # Sleep downtime (Seconds) (after simulated crash in SLEEP behavior)
    sleepDowntime = 8

}
```

### What happens in normal run?

All nodes start in follower behavior (some of them will change their behavior to candidate) and elect the first leader.

After some (configured) heartbeats from leader, the leader is simulating its crash and is "sleeping" for configured downtime. The next leader will be elected.

This happens again and again and again... until you stop the program or the earth is going to overheat. 😉

Data exchange (write data trough leader to followers) will be tested in test run (see below).

### What happens in test run?

1. Leader election (after init nodes)
2. Write data trough leader to followers (first write data to leader and replicate data to followers)
3. Get back data from all nodes (all nodes should have same data)
4. Simulate leader crash (triggered in test)
5. New leader election (old leader is gone)
6. Write data trough leader to followers (first write data to leader and replicate data to followers)
7. Get back data from all nodes (all nodes should have same data)


The ***integration-test*** is well documented - it is self explaining:

- ``./src/test/scala/de/maxbundscherer/scala/raft/RaftServiceTest.scala``

## Exciting (scala) stuff

Concurrent programming in Scala is usually done with akka actors. Akka actors is an actor model implementation for Scala and Java. Akka is developed/maintained by [Lightbend](https://www.lightbend.com/) (earlier called Typesafe).

The program and business logic is divided into separated actors. Each of these actors has its own state (own protected memory) and can only communicate with other actors by immutable messages.

![](./docImg/ActorModel.png)

([Image source](https://blog.scottlogic.com/2014/08/15/using-akka-and-scala-to-render-a-mandelbrot-set.html))

The ``RaftNodeActor`` has the following state implemented:

```scala
/**
    * Internal (mutable) actor state
    * @param neighbours Vector with another actors
    * @param electionTimer Cancellable for timer (used in FOLLOWER and CANDIDATE behavior)
    * @param heartbeatTimer Cancellable for timer (used in LEADER behavior)
    * @param alreadyVoted Boolean (has already voted in FOLLOWER behavior)
    * @param voteCounter Int (counter in CANDIDATE behavior)
    * @param majority Int (calculated majority - set up in init)
    * @param heartbeatCounter Int (auto simulate crash after some heartbeats in LEADER behavior)
   *  @param data Map (String->String) (used in FOLLOWER and LEADER behavior)
   *  @param lastHashCode Int (last hashcode from data) (used in FOLLOWER and LEADER behavior)
    */
  case class NodeState(
      var neighbours            : Vector[ActorRef]    = Vector.empty,
      var electionTimer         : Option[Cancellable] = None,
      var heartbeatTimer        : Option[Cancellable] = None,
      var alreadyVoted          : Boolean             = false,
      var voteCounter           : Int                 = 0,
      var majority              : Int                 = -1,
      var heartbeatCounter      : Int                 = 0,
      var data                  : Map[String, String] = Map.empty,
      var lastHashCode          : Int                 = -1,
  )
```

### Akka Actors Example

```scala
package de.maxbundscherer.scala.raft.examples

import akka.actor.{Actor, ActorLogging}

class SimpleActor extends Actor with ActorLogging {
  
  override def receive: Receive = {

    case data: String =>
      
      sender ! data + "-pong"

    case any: Any =>
      
      log.error(s"Got unhandled message '$any'")
      
  }
  
}
```

In this example, you can see a very simple akka actor: The actor is waiting for string-messages and replies with a new string (``!`` is used for [fire-and-forget-pattern](https://doc.akka.io/docs/akka/current/typed/interaction-patterns.html#fire-and-forget) / use ``?`` to use [ask-pattern](https://doc.akka.io/docs/akka/current/typed/interaction-patterns.html#request-response-with-ask-from-outside-an-actor) instead).

Non-string-messages are displayed by an error-logger.

### Raft nodes as akka actors

In this project, raft nodes are implemented as an akka actor (``RaftNodeActor``) with finite-state machine (FSM) behavior (see description and image above).

#### Finite-state machine (FSM) in akka

You can define multiple behaviors in an akka actor - see example:

```scala
package de.maxbundscherer.scala.raft.examples

import akka.actor.{Actor, ActorLogging}

object SimpleFSMActor {
  
  //Initialize message/command
  case class Initialize(state: Int)
  
}

class SimpleFSMActor extends Actor with ActorLogging {

  import SimpleFSMActor._
  
  //Actor mutable state
  private var state = -1

  //Initialized behavior 
  def initialized: Receive = {

    case any: Any => log.info(s"Got message '$any'")
    
  }

  //Default behavior
  override def receive: Receive = {

    case Initialize(newState) =>

      state = newState
      context.become(initialized)

    case any: Any => log.error(s"Not initialized '$any'")

  }

}
```

#### Service-Layer

Classic akka actors are not type safety. To "simulate" type safety, the service-layer (``RaftService``) was implemented. The service-layer is also used to spawn & initialize actors and to supervise the actor system - see examples:

- Spawn akka actor:
```scala
actorSystem.actorOf(props = RaftNodeActor.props, name = "myRaftNode")
```

- Ask (type safety non-blocking request):
```scala
def ping(): Future[Pong] = {
  ( actorRef ? Ping() ).asInstanceOf[Future[Pong]]
}
```

#### Aggregates

The object (read-only-singleton) ``RaftAggregate`` includes all necessary classes and objects (actor messages) for ``RaftService``, ``RaftNodeActor`` and ``RaftScheduler``.

#### Trait ``Configuration``

Scala traits are very similar to Java's interfaces. Traits can also include implementation. Normal classes can be extended (inheritance) by multiple traits, but only extend from one abstract class. Traits support multiple inheritance.

In this project the trait ``Configuration`` with internal object (read-only-singleton) ``Config`` is used to pass user-config to program.

The user-config is defined in the file ``application.conf`` and is loaded by a config-factory (see project dependencies).

#### Trait ``RaftScheduler``

The trait ``RaftScheduler`` is used to control raft-nodes timers in ``RaftNodeActor`` with the following function-calls:

- ``def stopElectionTimer()``: Used to stop electionTimer. This timer informs about "heartbeat-timeout" (``SchedulerTrigger.ElectionTimeout``) in FOLLOWER behavior and about "election-timeout" (``SchedulerTrigger.ElectionTimeout``) in CANDIDATE behavior.
- ``def restartElectionTimer()``: Used to stop and start electionTimer.
- ``def stopHeartbeatTimer()``: Used to stop heartbeatTimer. This timer informs about "send-heartbeat to all followers" (``SchedulerTrigger.Heartbeat``) in LEADER behavior.
- ``def restartHeartbeatTimer()``: Used to stop and start heartbeatTimer.
- ``def scheduleAwake()``: Used to trigger awakening automatically after downtime in SLEEP behavior (``SchedulerTrigger.Awake``). Awakening means: The node changes to follower-behavior.

Timers are controlled by ``changeBehavior`` and ``followerBehavior`` in ``RaftNodeActor`` to stop and start timers dependent on the nodes' behavior:

```scala
/**
   * Before change of behavior
   */
val newBehavior: Receive = toBehavior match {

  [...]
  
  case BehaviorEnum.FOLLOWER =>
    restartElectionTimer()
    stopHeartbeatTimer()
    followerBehavior

  case BehaviorEnum.CANDIDATE =>
    restartElectionTimer()
    stopHeartbeatTimer()
    candidateBehavior

  [...]

}
```

```scala
/**
   * After change of behavior
   */
toBehavior match {
      
  [...]

  case BehaviorEnum.SLEEP => scheduleAwake()
      
  [...]

}
```

```scala
/**
   * In followerBehavior
   */
case Heartbeat(lastHashCode) =>
      
      [...]

      restartElectionTimer()
```

#### Service Configurator Pattern

The program architecture is based on the [Service Configurator Pattern](https://www.usenix.org/legacy/publications/library/proceedings/coots97/full_papers/jain/jain.pdf).

The actor system & the services are started and configured in ...

- ... object ``Main`` for normal run.
- ... trait ``BaseServiceTest`` for test run.

## Scala compared to Go

- Data-types in Scala and Go are strong, static, inferred and structural typed.
- Scala intends to multicore architectures and brings functional programming & object oriented programming together. To improve code quality, you should not mix both concepts.
- Go intends to multicore architectures, too, and is an alternative to the programming language C.
- Learning Scala is time-consuming and sometimes quite involved because of the necessity to be familiar with the concept of functional programming and the huge amounts of complex concepts in the basic-language implemented.
- Learning Go is not so time-consuming, because Go is built on easy & familiar concepts (for example the concept of object oriented programming).
- Scala is usually running in the Java-virtual-machine (JVM) and can interact with Java-libraries. Compiling Scala [native](https://github.com/scala-native/scala-native) is possible, but unusual.
- Go is running native (is not compiled to byte-code) and can interact with C-libraries.

### Go concurrency

The language provides multiple possibilities:

- Concurrent execution ([goroutines](https://golangbot.com/goroutines/))
- Synchronization and messaging ([channels](https://www.geeksforgeeks.org/channel-in-golang/) - very similar to akka actors - Buffered Channels - FIFO)
- Multi-way concurrent control ([select](https://gobyexample.com/select))
- Low level blocking primitives ([locks/sync](https://golang.org/pkg/sync/))

### Scala concurrency

In Scala ``ExecutionContext`` (default is ``ExecutionContext.global``) is responsible for executing computations. The default ``ExecutionContext`` is a global static thread pool and is based on [Java's Fork/Join](https://docs.oracle.com/javase/tutorial/essential/concurrency/forkjoin.html). For example, you can set:

- ``scala.concurrent.context.minThreads``
- ``scala.concurrent.context.maxThreads``

You can also use multiple ``ExecutionContext``s in your application (or server-cluster).

Concurrent programming in Scala is usually done with akka actors. See "Exciting (scala) stuff" above.

You can also use:

- Scala Futures

```scala
val future = Future {
  getData()
}

future.onComplete {
  case Success(data) => println(s"Got $data")
  case Failure(exception) => println(s"Got failure $exception")
}
```

- Threads and Thread Pools from Java (unusually)

```scala
// This is unusually. Better use akka actors.

class ExampleProcessor extends Thread {
  override def run() {
    while(true) {
      val examples = Examples.getExamples()
      examples.foreach{ example =>
        process(example)
      }
    }
  }
}

//Start new thread
val thread = new ExampleProcessor()
thread.start()

//Wait for finishing
thread.join()
```

### My personal opinion:

- Scala is more empowering and you need less code.
- Go runs faster and very effective but sometimes feels repetitive and very mechanic.
- Scala is used for high-level cloud-applications (for example [Apache Spark](https://spark.apache.org/)).
- Go is used for low-level applications to make high-level-applications possible (for example [Docker](https://www.docker.com/)).
- **Comparing both languages is quite inconclusive because of their different fields of application.**

## Prospects

- This implementation is a prototype and should not be used in production.
- You can use [akka cluster](https://doc.akka.io/docs/akka/current/cluster-usage.html) to run this implementation on network and different machines. You have to modify the ``RaftService`` to spawn actors in cluster.
- Do not use Java serializer in production. It is slow and not secure. Use [Protobuf](https://github.com/protocolbuffers/protobuf) instead.


![](./docImg/logos.png)