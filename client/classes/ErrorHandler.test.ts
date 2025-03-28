import { test, expect } from 'bun:test';
import { ErrorHandler } from './ErrorHandler';
import { MissingEndpointError } from './errors/MissingEndpointError';
import { MalformedRequestError } from './errors/MalformedRequestError';

test('Throw missing endpoint', () => {
  expect(() => new ErrorHandler({
    apiVersion: 'v0.0.0'
  }).handleStatus(404)).toThrowError(MissingEndpointError);
});

test('Throw malformed request', () => {
  expect(() => new ErrorHandler({
    apiVersion: 'v0.0.0'
  }).handleStatus(400)).toThrowError(MalformedRequestError);
});
