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
    TypeDate,
    TypeDateFromJSON,
    TypeDateFromJSONTyped,
    TypeDateToJSON,
} from './';

/**
 * 
 * @export
 * @interface V1User
 */
export interface V1User {
    /**
     * 
     * @type {string}
     * @memberof V1User
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof V1User
     */
    firstName?: string;
    /**
     * 
     * @type {string}
     * @memberof V1User
     */
    lastName?: string;
    /**
     * 
     * @type {TypeDate}
     * @memberof V1User
     */
    birthday?: TypeDate;
    /**
     * 
     * @type {string}
     * @memberof V1User
     */
    age?: string;
}

export function V1UserFromJSON(json: any): V1User {
    return V1UserFromJSONTyped(json, false);
}

export function V1UserFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1User {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'firstName': !exists(json, 'firstName') ? undefined : json['firstName'],
        'lastName': !exists(json, 'lastName') ? undefined : json['lastName'],
        'birthday': !exists(json, 'birthday') ? undefined : TypeDateFromJSON(json['birthday']),
        'age': !exists(json, 'age') ? undefined : json['age'],
    };
}

export function V1UserToJSON(value?: V1User | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'firstName': value.firstName,
        'lastName': value.lastName,
        'birthday': TypeDateToJSON(value.birthday),
        'age': value.age,
    };
}

