declare let ClipboardJS: any;

// Apollo Client type overrides to fix Angular 12 build issues
declare module '@apollo/client/masking/internal/types' {
  type Exact<T> = { [P in keyof T]: T[P] };
}

declare module 'apollo-angular/query-ref' {
  interface UpdateQueryOptions<TVars> {
    variables?: TVars;
  }
}
