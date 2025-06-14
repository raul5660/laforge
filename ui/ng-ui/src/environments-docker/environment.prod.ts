// Production environment for Docker deployment

export const environment = {
  production: true,
  appVersion: 'v710',
  USERDATA_KEY: 'authf649fc9a5f55',
  isMockEnabled: false,
  apiUrl: 'api',
  graphqlUrl: '/api/query',
  wsUrl: 'ws://' + window.location.host + '/api/query',
  isMockApi: false,
  authBaseUrl: '/auth'
};
