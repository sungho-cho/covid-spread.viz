/**
 * @fileoverview gRPC-Web generated client stub for protos
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as covid_data_pb from './covid_data_pb';


export class CovidDataClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'binary';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorGetCountriesData = new grpcWeb.MethodDescriptor(
    '/protos.CovidData/GetCountriesData',
    grpcWeb.MethodType.UNARY,
    covid_data_pb.GetCountriesDataRequest,
    covid_data_pb.GetCountriesDataResponse,
    (request: covid_data_pb.GetCountriesDataRequest) => {
      return request.serializeBinary();
    },
    covid_data_pb.GetCountriesDataResponse.deserializeBinary
  );

  getCountriesData(
    request: covid_data_pb.GetCountriesDataRequest,
    metadata: grpcWeb.Metadata | null): Promise<covid_data_pb.GetCountriesDataResponse>;

  getCountriesData(
    request: covid_data_pb.GetCountriesDataRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: covid_data_pb.GetCountriesDataResponse) => void): grpcWeb.ClientReadableStream<covid_data_pb.GetCountriesDataResponse>;

  getCountriesData(
    request: covid_data_pb.GetCountriesDataRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: covid_data_pb.GetCountriesDataResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/protos.CovidData/GetCountriesData',
        request,
        metadata || {},
        this.methodDescriptorGetCountriesData,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/protos.CovidData/GetCountriesData',
    request,
    metadata || {},
    this.methodDescriptorGetCountriesData);
  }

  methodDescriptorGetMostRecentDate = new grpcWeb.MethodDescriptor(
    '/protos.CovidData/GetMostRecentDate',
    grpcWeb.MethodType.UNARY,
    covid_data_pb.Empty,
    covid_data_pb.Date,
    (request: covid_data_pb.Empty) => {
      return request.serializeBinary();
    },
    covid_data_pb.Date.deserializeBinary
  );

  getMostRecentDate(
    request: covid_data_pb.Empty,
    metadata: grpcWeb.Metadata | null): Promise<covid_data_pb.Date>;

  getMostRecentDate(
    request: covid_data_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: covid_data_pb.Date) => void): grpcWeb.ClientReadableStream<covid_data_pb.Date>;

  getMostRecentDate(
    request: covid_data_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: covid_data_pb.Date) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/protos.CovidData/GetMostRecentDate',
        request,
        metadata || {},
        this.methodDescriptorGetMostRecentDate,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/protos.CovidData/GetMostRecentDate',
    request,
    metadata || {},
    this.methodDescriptorGetMostRecentDate);
  }

  methodDescriptorGetAllData = new grpcWeb.MethodDescriptor(
    '/protos.CovidData/GetAllData',
    grpcWeb.MethodType.UNARY,
    covid_data_pb.Empty,
    covid_data_pb.GetAllDataResponse,
    (request: covid_data_pb.Empty) => {
      return request.serializeBinary();
    },
    covid_data_pb.GetAllDataResponse.deserializeBinary
  );

  getAllData(
    request: covid_data_pb.Empty,
    metadata: grpcWeb.Metadata | null): Promise<covid_data_pb.GetAllDataResponse>;

  getAllData(
    request: covid_data_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: covid_data_pb.GetAllDataResponse) => void): grpcWeb.ClientReadableStream<covid_data_pb.GetAllDataResponse>;

  getAllData(
    request: covid_data_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: covid_data_pb.GetAllDataResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/protos.CovidData/GetAllData',
        request,
        metadata || {},
        this.methodDescriptorGetAllData,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/protos.CovidData/GetAllData',
    request,
    metadata || {},
    this.methodDescriptorGetAllData);
  }

}

