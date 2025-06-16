#!/bin/bash

# Fix Apollo Client type errors by patching the problematic files
export NG_CLI_ANALYTICS="false"
export NODE_OPTIONS="--openssl-legacy-provider --max-old-space-size=4096"

# Patch the Apollo Client types file that's causing issues
APOLLO_CLIENT_TYPES="node_modules/@apollo/client/masking/internal/types.d.ts"

# Check if the file exists
if [ -f "$APOLLO_CLIENT_TYPES" ]; then
  echo "Patching Apollo Client types..."
  # Create a simplified version of the Apollo Client types file
  cat > "$APOLLO_CLIENT_TYPES" << 'EOL'
/* Patched version of Apollo Client types to work with Angular 15 / TypeScript 4.9 */
export interface ExactType<T> extends T { __exact?: never; }
export type Exact<T> = { [P in keyof T]: T[P] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: T[SubKey] };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: T[SubKey] | null };
export type ComputeRaw<T> = { [K in keyof T]: T[K] };
export type IntersectionOf<U> = {} & U;
export type Subtract<A, B> = Omit<A, keyof B> & {
  [K in keyof A & keyof B]?: undefined
};
EOL
  echo "Apollo Client types replaced successfully."
fi

# For Angular 15, we keep the standard polyfills.ts file

# This will run Angular build with production configuration
node_modules/.bin/ng build --configuration production
