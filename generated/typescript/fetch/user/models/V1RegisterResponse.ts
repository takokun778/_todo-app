/* tslint:disable */
/* eslint-disable */
/**
 * proto/proto/user/v1/user.proto
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
    V1User,
    V1UserFromJSON,
    V1UserFromJSONTyped,
    V1UserToJSON,
} from './';

/**
 * 
 * @export
 * @interface V1RegisterResponse
 */
export interface V1RegisterResponse {
    /**
     * 
     * @type {V1User}
     * @memberof V1RegisterResponse
     */
    user?: V1User;
}

export function V1RegisterResponseFromJSON(json: any): V1RegisterResponse {
    return V1RegisterResponseFromJSONTyped(json, false);
}

export function V1RegisterResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1RegisterResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'user': !exists(json, 'user') ? undefined : V1UserFromJSON(json['user']),
    };
}

export function V1RegisterResponseToJSON(value?: V1RegisterResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'user': V1UserToJSON(value.user),
    };
}

