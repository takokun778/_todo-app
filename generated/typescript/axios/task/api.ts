/* tslint:disable */
/* eslint-disable */
/**
 * proto/proto/task/v1/task.proto
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: version not set
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import { Configuration } from './configuration';
import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
// Some imports not used depending on template conditions
// @ts-ignore
import { DUMMY_BASE_URL, assertParamExists, setApiKeyToObject, setBasicAuthToObject, setBearerAuthToObject, setOAuthToObject, setSearchParams, serializeDataIfNeeded, toPathString, createRequestFunction } from './common';
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from './base';

/**
 * 
 * @export
 * @interface GooglerpcStatus
 */
export interface GooglerpcStatus {
    /**
     * 
     * @type {number}
     * @memberof GooglerpcStatus
     */
    code?: number;
    /**
     * 
     * @type {string}
     * @memberof GooglerpcStatus
     */
    message?: string;
    /**
     * 
     * @type {Array<ProtobufAny>}
     * @memberof GooglerpcStatus
     */
    details?: Array<ProtobufAny>;
}
/**
 * 
 * @export
 * @interface InlineObject
 */
export interface InlineObject {
    /**
     * 
     * @type {string}
     * @memberof InlineObject
     */
    title?: string;
    /**
     * 
     * @type {string}
     * @memberof InlineObject
     */
    detail?: string;
    /**
     * 
     * @type {string}
     * @memberof InlineObject
     */
    deadline?: string;
}
/**
 * 
 * @export
 * @interface ProtobufAny
 */
export interface ProtobufAny {
    /**
     * 
     * @type {string}
     * @memberof ProtobufAny
     */
    typeUrl?: string;
    /**
     * 
     * @type {string}
     * @memberof ProtobufAny
     */
    value?: string;
}
/**
 * 
 * @export
 * @interface V1CreateRequest
 */
export interface V1CreateRequest {
    /**
     * 
     * @type {string}
     * @memberof V1CreateRequest
     */
    creatorId?: string;
    /**
     * 
     * @type {string}
     * @memberof V1CreateRequest
     */
    title?: string;
    /**
     * 
     * @type {string}
     * @memberof V1CreateRequest
     */
    detail?: string;
    /**
     * 
     * @type {string}
     * @memberof V1CreateRequest
     */
    deadline?: string;
}
/**
 * 
 * @export
 * @interface V1CreateResponse
 */
export interface V1CreateResponse {
    /**
     * 
     * @type {V1Task}
     * @memberof V1CreateResponse
     */
    task?: V1Task;
}
/**
 * 
 * @export
 * @interface V1DeleteResponse
 */
export interface V1DeleteResponse {
    /**
     * 
     * @type {V1Task}
     * @memberof V1DeleteResponse
     */
    task?: V1Task;
}
/**
 * 
 * @export
 * @interface V1DoneResponse
 */
export interface V1DoneResponse {
    /**
     * 
     * @type {V1Task}
     * @memberof V1DoneResponse
     */
    task?: V1Task;
}
/**
 * 
 * @export
 * @interface V1GetResponse
 */
export interface V1GetResponse {
    /**
     * 
     * @type {V1Task}
     * @memberof V1GetResponse
     */
    task?: V1Task;
}
/**
 * 
 * @export
 * @interface V1ListResponse
 */
export interface V1ListResponse {
    /**
     * 
     * @type {Array<V1Task>}
     * @memberof V1ListResponse
     */
    task?: Array<V1Task>;
}
/**
 * 
 * @export
 * @interface V1Task
 */
export interface V1Task {
    /**
     * 
     * @type {string}
     * @memberof V1Task
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Task
     */
    creatorId?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Task
     */
    title?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Task
     */
    detail?: string;
    /**
     * 
     * @type {V1TaskStatus}
     * @memberof V1Task
     */
    status?: V1TaskStatus;
    /**
     * 
     * @type {string}
     * @memberof V1Task
     */
    createdAt?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Task
     */
    updatedAt?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Task
     */
    deadline?: string;
}
/**
 * 
 * @export
 * @enum {string}
 */

export enum V1TaskStatus {
    Incomplete = 'INCOMPLETE',
    Complete = 'COMPLETE'
}

/**
 * 
 * @export
 * @interface V1UndoneResponse
 */
export interface V1UndoneResponse {
    /**
     * 
     * @type {V1Task}
     * @memberof V1UndoneResponse
     */
    task?: V1Task;
}
/**
 * 
 * @export
 * @interface V1UpdateResponse
 */
export interface V1UpdateResponse {
    /**
     * 
     * @type {V1Task}
     * @memberof V1UpdateResponse
     */
    task?: V1Task;
}

/**
 * TaskServiceApi - axios parameter creator
 * @export
 */
export const TaskServiceApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @param {V1CreateRequest} body 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        taskServiceCreate: async (body: V1CreateRequest, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'body' is not null or undefined
            assertParamExists('taskServiceCreate', 'body', body)
            const localVarPath = `/task`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter, options.query);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(body, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @param {string} id 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        taskServiceDelete: async (id: string, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            assertParamExists('taskServiceDelete', 'id', id)
            const localVarPath = `/task/{id}`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'DELETE', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter, options.query);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @param {string} id 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        taskServiceDone: async (id: string, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            assertParamExists('taskServiceDone', 'id', id)
            const localVarPath = `/task/{id}/done`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'PUT', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter, options.query);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @param {string} id 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        taskServiceGet: async (id: string, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            assertParamExists('taskServiceGet', 'id', id)
            const localVarPath = `/task/{id}`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter, options.query);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @param {string} userId 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        taskServiceList: async (userId: string, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'userId' is not null or undefined
            assertParamExists('taskServiceList', 'userId', userId)
            const localVarPath = `/tasks/{userId}`
                .replace(`{${"userId"}}`, encodeURIComponent(String(userId)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter, options.query);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @param {string} id 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        taskServiceUndone: async (id: string, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            assertParamExists('taskServiceUndone', 'id', id)
            const localVarPath = `/task/{id}/undone`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'PUT', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter, options.query);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @param {string} id 
         * @param {InlineObject} body 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        taskServiceUpdate: async (id: string, body: InlineObject, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            assertParamExists('taskServiceUpdate', 'id', id)
            // verify required parameter 'body' is not null or undefined
            assertParamExists('taskServiceUpdate', 'body', body)
            const localVarPath = `/task/{id}`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'PUT', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter, options.query);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(body, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * TaskServiceApi - functional programming interface
 * @export
 */
export const TaskServiceApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = TaskServiceApiAxiosParamCreator(configuration)
    return {
        /**
         * 
         * @param {V1CreateRequest} body 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async taskServiceCreate(body: V1CreateRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<V1CreateResponse>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.taskServiceCreate(body, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @param {string} id 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async taskServiceDelete(id: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<V1DeleteResponse>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.taskServiceDelete(id, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @param {string} id 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async taskServiceDone(id: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<V1DoneResponse>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.taskServiceDone(id, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @param {string} id 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async taskServiceGet(id: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<V1GetResponse>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.taskServiceGet(id, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @param {string} userId 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async taskServiceList(userId: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<V1ListResponse>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.taskServiceList(userId, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @param {string} id 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async taskServiceUndone(id: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<V1UndoneResponse>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.taskServiceUndone(id, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @param {string} id 
         * @param {InlineObject} body 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async taskServiceUpdate(id: string, body: InlineObject, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<V1UpdateResponse>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.taskServiceUpdate(id, body, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
    }
};

/**
 * TaskServiceApi - factory interface
 * @export
 */
export const TaskServiceApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = TaskServiceApiFp(configuration)
    return {
        /**
         * 
         * @param {V1CreateRequest} body 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        taskServiceCreate(body: V1CreateRequest, options?: any): AxiosPromise<V1CreateResponse> {
            return localVarFp.taskServiceCreate(body, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @param {string} id 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        taskServiceDelete(id: string, options?: any): AxiosPromise<V1DeleteResponse> {
            return localVarFp.taskServiceDelete(id, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @param {string} id 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        taskServiceDone(id: string, options?: any): AxiosPromise<V1DoneResponse> {
            return localVarFp.taskServiceDone(id, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @param {string} id 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        taskServiceGet(id: string, options?: any): AxiosPromise<V1GetResponse> {
            return localVarFp.taskServiceGet(id, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @param {string} userId 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        taskServiceList(userId: string, options?: any): AxiosPromise<V1ListResponse> {
            return localVarFp.taskServiceList(userId, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @param {string} id 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        taskServiceUndone(id: string, options?: any): AxiosPromise<V1UndoneResponse> {
            return localVarFp.taskServiceUndone(id, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @param {string} id 
         * @param {InlineObject} body 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        taskServiceUpdate(id: string, body: InlineObject, options?: any): AxiosPromise<V1UpdateResponse> {
            return localVarFp.taskServiceUpdate(id, body, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * TaskServiceApi - object-oriented interface
 * @export
 * @class TaskServiceApi
 * @extends {BaseAPI}
 */
export class TaskServiceApi extends BaseAPI {
    /**
     * 
     * @param {V1CreateRequest} body 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof TaskServiceApi
     */
    public taskServiceCreate(body: V1CreateRequest, options?: any) {
        return TaskServiceApiFp(this.configuration).taskServiceCreate(body, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @param {string} id 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof TaskServiceApi
     */
    public taskServiceDelete(id: string, options?: any) {
        return TaskServiceApiFp(this.configuration).taskServiceDelete(id, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @param {string} id 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof TaskServiceApi
     */
    public taskServiceDone(id: string, options?: any) {
        return TaskServiceApiFp(this.configuration).taskServiceDone(id, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @param {string} id 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof TaskServiceApi
     */
    public taskServiceGet(id: string, options?: any) {
        return TaskServiceApiFp(this.configuration).taskServiceGet(id, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @param {string} userId 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof TaskServiceApi
     */
    public taskServiceList(userId: string, options?: any) {
        return TaskServiceApiFp(this.configuration).taskServiceList(userId, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @param {string} id 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof TaskServiceApi
     */
    public taskServiceUndone(id: string, options?: any) {
        return TaskServiceApiFp(this.configuration).taskServiceUndone(id, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @param {string} id 
     * @param {InlineObject} body 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof TaskServiceApi
     */
    public taskServiceUpdate(id: string, body: InlineObject, options?: any) {
        return TaskServiceApiFp(this.configuration).taskServiceUpdate(id, body, options).then((request) => request(this.axios, this.basePath));
    }
}


