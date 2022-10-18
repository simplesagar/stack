import { TransactionRequest } from 'numary/dist/schema';
import { ledger } from './ledger';
import uniqid from 'uniqid';
import { program } from 'commander';

program.option('-l, --ledger <string>');
program.option('-r, --rounds <number>');
program.parse();
const options = program.opts();

const trade = (wallet: string, amountFiat: number, asset: string, amountAsset: number) : any => {
    const tradeId = uniqid().toUpperCase();

    return {
        postings: [
            {
                amount: amountFiat,
                asset: 'EUR/2',
                source: wallet,
                destination: `trades:${tradeId}`,
            },
            {
                amount: amountFiat,
                asset: 'EUR/2',
                source: `trades:${tradeId}`,
                destination: 'shares:fiat:holdings',
            },
            {
                amount: amountAsset,
                asset,
                source: 'teller:otc:nyse',
                destination: `trades:${tradeId}`,
            },
            {
                amount: amountAsset,
                asset,
                source: `trades:${tradeId}`,
                destination: wallet,
            },
        ],
    };
};

const populate = async () => {
    for (let userId = 1; userId <= (parseInt(options.rounds) || 1e3); userId++) {
        const txs : TransactionRequest[] = [];

        const wallet = `users:${userId}:wallet`;

        // deposit

        const id = uniqid().toUpperCase();

        txs.push({
            postings: [
                {
                    amount: 100e2,
                    asset: 'EUR/2',
                    source: `world`,
                    destination: `payments:adyen:${id}`,
                },
            ],
        });

        txs.push({
            postings: [
                {
                    amount: 100e2,
                    asset: 'EUR/2',
                    source: `payments:adyen:${id}`,
                    destination: wallet,
                }
            ],
        });

        // trade

        txs.push({
            postings: [
                {
                    amount: 0.35e6,
                    asset: 'RBLX/6',
                    source: 'world',
                    destination: 'teller:otc:nyse',
                },
                {
                    amount: 1.84e6,
                    asset: 'SNAP/6',
                    source: 'world',
                    destination: 'teller:otc:nyse',
                }
            ],
        });

        txs.push(trade(wallet, 15e2, 'RBLX/6', 0.35e6));
        txs.push(trade(wallet, 42.3e2, 'SNAP/6', 1.84e6));

        // withdrawal

        const withdrawal = `users:${userId}:withdrawals:${uniqid()}`;

        txs.push({
            postings: [
                {
                    amount: 22.7e2,
                    asset: 'EUR/2',
                    source: wallet,
                    destination: withdrawal,
                },
            ],
        });

        txs.push({
            postings: [
                {
                    amount: 22.7e2,
                    asset: 'EUR/2',
                    source: withdrawal,
                    destination: `payments:${uniqid()}`,
                },
            ],
        });

        const res = await ledger(options.ledger || 'investing-app-demo-11').batch(txs);

        console.log(res);
    }
}

(async() => {
    const table = []
    for (let i = 0; i < 100; i++) {
        const req = populate()
        table.push(req)
    }

    for (const req of table) {
        await req
    }
})();