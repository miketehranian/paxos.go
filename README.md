Simple example of Paxos written in go

Nomenclature used in Paxos:
Replicas - Processes


Each process below is its own DSM

Replicas receive requests from clients and asks leaders to serialize the requests so all replicas have the same sequence. Replicas also respond to clients. Applies the serialized sequence determined by the leader to the shared data.

Acceptor stores the fault tolerant memory for the acceptor.

Leaders serialize requests sent to the replicas.


Replicas has a sequence of slots that are indexed. When a request comes in it is called a proposal. The synod protocol is used to determine the single command to use for each slot, and awaits the decision from the leaders and acceptors before sending a response back to the client.


Ballots are handled by only one leader at a time. One ballot can handle multiple slots, but they are not slots.

Leaders send message to accepters; leaders are clients and acceptors are servers.


