import { NgModule } from '@angular/core';
import { ApolloClientOptions, InMemoryCache, split, ApolloLink } from '@apollo/client/core';
import { getMainDefinition } from '@apollo/client/utilities';
import { APOLLO_OPTIONS, ApolloModule } from 'apollo-angular';
import { HttpLink } from 'apollo-angular/http';
import { environment } from 'src/environments/environment';

// For WebSockets in Apollo Client 3.x
import { WebSocketLink } from '@apollo/client/link/ws';

export function createApollo(httpLink: HttpLink): ApolloClientOptions<any> {
  const httpClient = httpLink.create({
    uri: environment.graphqlUrl,
    withCredentials: true
  });
  // Using WebSocketLink which is compatible with Apollo Angular
  const wsClient = new WebSocketLink({
    uri: environment.wsUrl,
    options: {
      reconnect: true,
      connectionParams: {}
    }
  });

  const link = split(
    ({ query }) => {
      const { kind, operation } = getMainDefinition(query) as any;
      return kind === 'OperationDefinition' && operation === 'subscription';
    },
    wsClient,
    httpClient
  );

  return {
    uri: environment.graphqlUrl,
    link,
    cache: new InMemoryCache({
      resultCaching: false
    }),
    credentials: 'include'
  };
}

@NgModule({
  imports: [
    ApolloModule
  ],
  providers: [
    {
      provide: APOLLO_OPTIONS,
      useFactory: createApollo,
      deps: [HttpLink]
    }
  ]
})
export class GraphQLModule {}
