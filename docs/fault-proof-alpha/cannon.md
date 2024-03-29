## Generate Traces with `cannon` and `inura-program`

Normally, `inura-challenger` handles creating the required traces as part of responding to games. However, for manual
testing it may be useful to manually generate the trace. This can be done by running `cannon` directly.

### Prerequisites

- The cannon pre-state downloaded from [Goerli deployment](./deployments.md#goerli).
- A Goerli L1 node.
    - An archive node is not required.
    - Public RPC providers can be used, however a significant number of requests will need to be made which may exceed
      rate limits for free plans.
- An OP-Goerli L2 archive node with `debug` APIs enabled.
    - An archive node is required to ensure world-state pre-images remain available.
    - Public RPC providers are generally not usable as they don’t support the `debug_dbGet` RPC method.

### Compilation

To compile the required programs, in the top level of the monorepo run:

```bash
make cannon-prestate
```

This will compile the `cannon` executable to `cannon/bin/cannon` as well as the `inura-program` executable used to fetch
pre-image data to `inura-program/bin/inura-program`.

### Run Cannon

To run cannon to generate a proof use:

```bash
mkdir -p temp/cannon/proofs temp/cannon/snapshots temp/cannon/preimages

./cannon/bin/cannon run \
    --pprof.cpu \
    --info-at '%10000000' \
    --proof-at '=<TRACE_INDEX>' \
    --stop-at '=<STOP_INDEX>' \
    --proof-fmt 'temp/cannon/proofs/%d.json' \
    --snapshot-at '%1000000000' \
    --snapshot-fmt 'temp/cannon/snapshots/%d.json.gz' \
    --input <PRESTATE> \
    --output temp/cannon/stop-state.json \
    -- \
    ./inura-program/bin/inura-program \
    --network goerli \
    --l1 <L1_URL> \
    --l2 <L2_URL> \
    --l1.head <L1_HEAD> \
    --l2.claim <L2_CLAIM> \
    --l2.head <L2_HEAD> \
    --l2.blocknumber <L2_BLOCK_NUMBER> \
    --datadir temp/cannon/preimages \
    --log.format terminal \
    --server
```

The placeholders are:

- `<TRACE_INDEX>` the index in the trace to generate a proof for
- `<STOP_INDEX>` the index to stop execution at. Typically this is one instruction after `<TRACE_INDEX>` to stop as soon
  as the required proof has been generated.
- `<PRESTATE>` the prestate.json downloaded above. Note that this needs to precisely match the prestate used on-chain so
  must be the downloaded version and not a version built locally.
- `<L1_URL>` the Goerli L1 JSON RPC endpoint
- `<L2_URL>` the OP-Goerli L2 archive node JSON RPC endpoint
- `<L1_HEAD>` the hash of the L1 head block used for the dispute game
- `<L2_CLAIM>` the output root immediately prior to the disputed root in the L2 output oracle
- `<L2_HEAD>` the hash of the L2 block that `<L2_CLAIM>`is from
- `<L2_BLOCK_NUMBER>` the block number that `<L2_CLAIM>` is from

The generated proof will be stored in the `temp/cannon/proofs/` directory. The hash to use as the claim value is
the `post` field of the generated proof which provides the hash of the cannon state witness after execution of the step.

Since cannon can be very slow to execute, the above command uses the `--snapshot-at` option to generate a snapshot of
the cannon state every 1000000000 instructions. Once generated, these snapshots can be used as the `--input` to begin
execution at that step rather than from the very beginning. Generated snapshots are stored in
the `temp/cannon/snapshots` directory.

See `./cannon/bin/cannon --help` for further information on the options available.

### Trace Extension

Fault dispute games always use a trace with a fixed length of `2 ^ MAX_GAME_DEPTH`. The trace generated by `cannon`
stops when the client program exits, so this trace must be extended by repeating the hash of the final state in the
actual trace for all remaining steps. Cannon does not perform this trace extension automatically.

If cannon stops execution before the trace index you requested a proof at, it simply will not generate a proof. When it
stops executing, it will write its final state to `temp/cannon/stop-state.json` (controlled by the `--output` option).
The `step` field of this state contains the last step cannon executed. Once the final step is known, rerun cannon to
generate the proof at that final step and use the `post` hash as the claim value for all later trace indices.
