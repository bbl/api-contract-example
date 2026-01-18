import { defineConfig } from '@hey-api/openapi-ts';

export default defineConfig({
  input: './generated/openapi.yaml',
  output: './generated/ts-client',
  plugins: [
    '@hey-api/typescript',
    {
      name: '@hey-api/sdk',
      operations: {
        strategy: 'single',
        containerName: 'ApiClient',
      },
    },
  ],
});
