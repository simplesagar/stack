import Cluster from "numary";
import { Cloud } from "numary/dist/cloud";
import Ledger from "numary/dist/ledger";

const cluster = new Cluster({
    uri: 'http://127.0.0.1:3068',
});

export const ledger = (name: string) : Ledger => {
    return cluster.getLedger(name);
}