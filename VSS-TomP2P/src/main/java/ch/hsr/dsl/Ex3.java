package ch.hsr.dsl;

import net.tomp2p.dht.FutureGet;
import net.tomp2p.dht.FuturePut;
import net.tomp2p.dht.PeerBuilderDHT;
import net.tomp2p.dht.PeerDHT;
import net.tomp2p.p2p.Peer;
import net.tomp2p.p2p.PeerBuilder;
import net.tomp2p.peers.Number160;
import net.tomp2p.storage.Data;

import java.io.IOException;

public class Ex3 {
    public static void main(String[] args) throws IOException, ClassNotFoundException {
        PeerDHT[] peers = null;
        PeerDHT[] peers2 = null;
        try {
            peers = ExampleUtils.createAndAttachPeersDHT(100, 4001);
            ExampleUtils.bootstrap(peers);
            //
            String data = "This is a data item";

            // put/get here

            //System.out.println("got 1: "+ fg.data().object());

            peers2 = new PeerDHT[peers.length + 5];
            for(int i=0;i<peers.length;i++) {
                peers2[i] = peers[i];
            }
            //lets insert a few peers close to the key hash(data), return the string attack

            // attack here

            //add attacker to the network
            ExampleUtils.bootstrap(peers2);

            //fg = peers[10].get(Number160.createHash(data)).start().awaitUninterruptibly();
            //System.out.println("got 2: "+ fg.data().object());

        } finally {
            // 0 is the master
            if (peers2 != null) {
                for(int i=0;i<peers2.length;i++) {
                    peers2[i].shutdown();
                }
            }
        }
    }
}
