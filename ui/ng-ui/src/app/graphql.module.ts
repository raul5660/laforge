import { NgModule } from '@angular/core';
import { ApolloModule, APOLLO_OPTIONS } from 'apollo-angular';
import { ApolloClientOptions, InMemoryCache, ApolloLink, split, NormalizedCacheObject, FetchResult } from '@apollo/client/core';
import { HttpLink } from 'apollo-angular/http';
import { Operation } from '@apollo/client/core';
import { environment } from 'src/environments/environment';
import { onError } from '@apollo/client/link/error';
import { print } from 'graphql';
import { Observable } from 'rxjs';
import { from } from 'rxjs';

// Import graphql-ws
import { createClient } from 'graphql-ws';

// Import helper to determine operation type
import { getMainDefinition } from '@apollo/client/utilities';

/**
 * This function creates an Apollo configuration with HTTP and WebSocket support
 * compatible with Angular 16, Apollo Angular 5.0 and graphql-ws
 */
export function createApollo(httpLink: HttpLink): ApolloClientOptions<NormalizedCacheObject> {
  console.log('Initializing GraphQL with Apollo Angular 5.0');
  console.log(`HTTP endpoint: ${environment.graphqlUrl}`);
  console.log(`WebSocket endpoint: ${environment.wsUrl}`);
  
  // Setup error handling with detailed logging
  const errorLink = onError(({ graphQLErrors, networkError, operation, response }) => {
    if (graphQLErrors) {
      graphQLErrors.forEach(({ message, locations, path }) => {
        console.error(`[GraphQL error]: Message: ${message}, Location: ${JSON.stringify(locations)}, Path: ${path}`);
        console.error(`Operation: ${operation.operationName || 'unnamed'}, variables: ${JSON.stringify(operation.variables)}`);
      });
    }

    if (networkError) {
      console.error(`[Network error for ${operation.operationName || 'unnamed operation'}]:`, networkError);
    }
    
    // Log full response for debugging
    console.log(`[Apollo ${operation.operationName || 'unnamed'} response]:`, response);
  });

  // Setup HTTP link for queries and mutations with credentials
  const http = httpLink.create({
    uri: environment.graphqlUrl,
    withCredentials: true // Important for cookie-based auth
  });

  // Generate dynamic WebSocket URL based on current window location
  const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
  const wsUrl = `${protocol}//${window.location.host}/api/query`;
  
  console.log(`Dynamic WebSocket URL: ${wsUrl}`);
  
  // Setup WebSocket client for subscriptions with enhanced debug handlers
  console.log('Initializing GraphQL WebSocket client at:', wsUrl);
  const wsClient = createClient({
    url: wsUrl,
    connectionParams: {},
    retryAttempts: 5,
    keepAlive: 10000,
    on: {
      connecting: () => console.log('WebSocket connecting...'),
      connected: () => {
        console.log('WebSocket connected successfully');
        // Dispatch an event that other components can listen for
        window.dispatchEvent(new CustomEvent('websocket-connected'));
      },
      error: (e) => {
        console.error('WebSocket connection error:', e);
        // Log more details about the error
        if (e) {
          const errorDetails: Record<string, unknown> = {};
          if (typeof e === 'object') {
            // Type guard for common error properties
            const err = e as any;
            if (err.message) errorDetails.message = err.message;
            if (err.type) errorDetails.type = err.type;
            if (err.code) errorDetails.code = err.code;
            if (err.target) errorDetails.target = String(err.target);
          }
          console.error('Connection error details:', errorDetails);
        }
      },
      closed: (event) => {
        if (event && typeof event === 'object') {
          const evt = event as any;
          console.log('WebSocket connection closed with code:', evt.code || 'unknown');
          console.log('Close reason:', evt.reason || 'no reason provided');
        } else {
          console.log('WebSocket connection closed (no details available)');
        }
      },
      message: (message) => {
        if (message && typeof message === 'object') {
          const msg = message as any;
          console.debug('WebSocket message received', msg.type || 'unknown type');
        } else {
          console.debug('WebSocket message received (no details available)');
        }
      }
    }
  });

  // Create a custom WebSocket link
  const ws = new ApolloLink((operation: Operation): any => {
    return new Observable<FetchResult>((observer) => {
      try {
        const sub = wsClient.subscribe(
          {
            query: print(operation.query),
            variables: operation.variables,
            operationName: operation.operationName
          },
          {
            next: (result) => {
              observer.next(result);
            },
            error: (error) => {
              observer.error(error);
            },
            complete: () => {
              observer.complete();
            }
          }
        );

        return () => sub();
      } catch (err) {
        observer.error(err);
        return () => {};
      }
    });
  });

  // Use split to route subscription operations to WebSocket and others to HTTP
  const link = split(
    ({ query }) => {
      const definition = getMainDefinition(query);
      return (
        definition.kind === 'OperationDefinition' && 
        definition.operation === 'subscription'
      );
    },
    ws,
    errorLink.concat(http)
  );


  return {
    link,
    cache: new InMemoryCache({
      resultCaching: false
    }),
    defaultOptions: {
      watchQuery: {
        fetchPolicy: 'network-only',
        errorPolicy: 'all'
      },
      query: {
        fetchPolicy: 'network-only',
        errorPolicy: 'all'
      },
      mutate: {
        errorPolicy: 'all'
      }
    }
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
