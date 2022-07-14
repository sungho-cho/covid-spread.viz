import * as jspb from 'google-protobuf'



export class Date extends jspb.Message {
  getYear(): number;
  setYear(value: number): Date;

  getMonth(): number;
  setMonth(value: number): Date;

  getDay(): number;
  setDay(value: number): Date;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Date.AsObject;
  static toObject(includeInstance: boolean, msg: Date): Date.AsObject;
  static serializeBinaryToWriter(message: Date, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Date;
  static deserializeBinaryFromReader(message: Date, reader: jspb.BinaryReader): Date;
}

export namespace Date {
  export type AsObject = {
    year: number,
    month: number,
    day: number,
  }
}

export class GetAllDataResponse extends jspb.Message {
  getFirstDate(): Date | undefined;
  setFirstDate(value?: Date): GetAllDataResponse;
  hasFirstDate(): boolean;
  clearFirstDate(): GetAllDataResponse;

  getLastDate(): Date | undefined;
  setLastDate(value?: Date): GetAllDataResponse;
  hasLastDate(): boolean;
  clearLastDate(): GetAllDataResponse;

  getDataList(): Array<CountriesData>;
  setDataList(value: Array<CountriesData>): GetAllDataResponse;
  clearDataList(): GetAllDataResponse;
  addData(value?: CountriesData, index?: number): CountriesData;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAllDataResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetAllDataResponse): GetAllDataResponse.AsObject;
  static serializeBinaryToWriter(message: GetAllDataResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAllDataResponse;
  static deserializeBinaryFromReader(message: GetAllDataResponse, reader: jspb.BinaryReader): GetAllDataResponse;
}

export namespace GetAllDataResponse {
  export type AsObject = {
    firstDate?: Date.AsObject,
    lastDate?: Date.AsObject,
    dataList: Array<CountriesData.AsObject>,
  }
}

export class GetCountriesDataRequest extends jspb.Message {
  getDate(): Date | undefined;
  setDate(value?: Date): GetCountriesDataRequest;
  hasDate(): boolean;
  clearDate(): GetCountriesDataRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetCountriesDataRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetCountriesDataRequest): GetCountriesDataRequest.AsObject;
  static serializeBinaryToWriter(message: GetCountriesDataRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetCountriesDataRequest;
  static deserializeBinaryFromReader(message: GetCountriesDataRequest, reader: jspb.BinaryReader): GetCountriesDataRequest;
}

export namespace GetCountriesDataRequest {
  export type AsObject = {
    date?: Date.AsObject,
  }
}

export class GetCountriesDataResponse extends jspb.Message {
  getCountriesData(): CountriesData | undefined;
  setCountriesData(value?: CountriesData): GetCountriesDataResponse;
  hasCountriesData(): boolean;
  clearCountriesData(): GetCountriesDataResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetCountriesDataResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetCountriesDataResponse): GetCountriesDataResponse.AsObject;
  static serializeBinaryToWriter(message: GetCountriesDataResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetCountriesDataResponse;
  static deserializeBinaryFromReader(message: GetCountriesDataResponse, reader: jspb.BinaryReader): GetCountriesDataResponse;
}

export namespace GetCountriesDataResponse {
  export type AsObject = {
    countriesData?: CountriesData.AsObject,
  }
}

export class CountryData extends jspb.Message {
  getCountry(): string;
  setCountry(value: string): CountryData;

  getIso3sList(): Array<string>;
  setIso3sList(value: Array<string>): CountryData;
  clearIso3sList(): CountryData;
  addIso3s(value: string, index?: number): CountryData;

  getDate(): Date | undefined;
  setDate(value?: Date): CountryData;
  hasDate(): boolean;
  clearDate(): CountryData;

  getConfirmed(): number;
  setConfirmed(value: number): CountryData;

  getRecovered(): number;
  setRecovered(value: number): CountryData;

  getDeaths(): number;
  setDeaths(value: number): CountryData;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CountryData.AsObject;
  static toObject(includeInstance: boolean, msg: CountryData): CountryData.AsObject;
  static serializeBinaryToWriter(message: CountryData, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CountryData;
  static deserializeBinaryFromReader(message: CountryData, reader: jspb.BinaryReader): CountryData;
}

export namespace CountryData {
  export type AsObject = {
    country: string,
    iso3sList: Array<string>,
    date?: Date.AsObject,
    confirmed: number,
    recovered: number,
    deaths: number,
  }
}

export class CountriesData extends jspb.Message {
  getDate(): Date | undefined;
  setDate(value?: Date): CountriesData;
  hasDate(): boolean;
  clearDate(): CountriesData;

  getCountriesList(): Array<CountryData>;
  setCountriesList(value: Array<CountryData>): CountriesData;
  clearCountriesList(): CountriesData;
  addCountries(value?: CountryData, index?: number): CountryData;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CountriesData.AsObject;
  static toObject(includeInstance: boolean, msg: CountriesData): CountriesData.AsObject;
  static serializeBinaryToWriter(message: CountriesData, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CountriesData;
  static deserializeBinaryFromReader(message: CountriesData, reader: jspb.BinaryReader): CountriesData;
}

export namespace CountriesData {
  export type AsObject = {
    date?: Date.AsObject,
    countriesList: Array<CountryData.AsObject>,
  }
}

export class Empty extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Empty.AsObject;
  static toObject(includeInstance: boolean, msg: Empty): Empty.AsObject;
  static serializeBinaryToWriter(message: Empty, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Empty;
  static deserializeBinaryFromReader(message: Empty, reader: jspb.BinaryReader): Empty;
}

export namespace Empty {
  export type AsObject = {
  }
}

