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

export function V1CreateResponseFromJSON(json: any): V1CreateResponse {
    return V1CreateResponseFromJSONTyped(json, false);
}

export function V1CreateResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1CreateResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'task': !exists(json, 'task') ? undefined : V1TaskFromJSON(json['task']),
    };
}

export function V1CreateResponseToJSON(value?: V1CreateResponse | null): any {
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


