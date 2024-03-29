## Contracts TS

[![codecov](https://codecov.io/gh/inuraorg/inura/branch/develop/graph/badge.svg?token=0VTG7PG7YR&flag=contracts-bedrock-tests)](https://codecov.io/gh/inuraorg/inura)

ABI and Address constants + generated code from [@eth-inura/contracts-bedrock/](../contracts-bedrock/) for use in TypeScript.

Much of this package is generated. See [CODE_GEN.md](./CODE_GEN.md) for instructions on how to generate.

#### @eth-inura/contracts-ts

The main entrypoint exports constants related to contracts bedrock as const. As const allows it to be [used in TypeScript with stronger typing than importing JSON](https://github.com/microsoft/TypeScript/issues/32063).

- Exports contract abis.
- Exports contract addresses

```typescript
import {
  l2OutputOracleProxyABI,
  l2OutputOracleAddresses,
} from '@eth-inura/contracts-ts'

console.log(l2OutputOracleAddresses[10], abi)
```

Addresses are also exported as an object for convenience.

```typescript
import { addresses } from '@eth-inura/contracts-ts'

console.log(addresses.l2OutputOracle[10])
```

#### @eth-inura/contracts-ts/react

- All [React hooks](https://wagmi.sh/cli/plugins/react) `@eth-inura/contracts-ts/react`

```typescript
import { useAddressManagerAddress } from '@eth-inura/contracts-ts/react'

const component = () => {
  const { data, error, loading } = useAddressManagerAddress()
  if (loading) {
    return <div>Loading</div>
  }
  if (err) {
    return <div>Error</div>
  }
  return <div>{data}</div>
}
```

#### @eth-inura/contracts-ts/actions

- All [wagmi actions](https://wagmi.sh/react/actions) for use in Vanilla JS or non react code

```typescript
import { readSystemConfig } from '@eth-inura/contracts-ts/actions'
console.log(await readSystemConfig())
```

#### See Also

- [Contracts bedrock specs](../../specs/)
- [Wagmi](https://wagmi.sh)
