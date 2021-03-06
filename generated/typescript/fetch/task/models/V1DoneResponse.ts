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

import { exists, mapValues } from '../runtime';
import {
    V1Task,
    V1TaskFromJSON,
    V1TaskFromJSONTyped,
    V1TaskToJSON,
} from './';

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

export function V1DoneResponseFromJSON(json: any): V1DoneResponse {
    return V1DoneResponseFromJSONTyped(json, false);
}

export function V1DoneResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1DoneResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'task': !exists(json, 'task') ? undefined : V1TaskFromJSON(json['task']),
    };
}

export function V1DoneResponseToJSON(value?: V1DoneResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'task': V1TaskToJSON(value.task),
    };
}


