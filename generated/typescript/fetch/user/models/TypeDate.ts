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
/**
 * * A full date, with non-zero year, month, and day values
 * * A month and day value, with a zero year, such as an anniversary
 * * A year on its own, with zero month and day values
 * * A year and month value, with a zero day, such as a credit card expiration
 * date
 * 
 * Related types are [google.type.TimeOfDay][google.type.TimeOfDay] and
 * `google.protobuf.Timestamp`.
 * @export
 * @interface TypeDate
 */
export interface TypeDate {
    /**
     * Year of the date. Must be from 1 to 9999, or 0 to specify a date without
     * a year.
     * @type {number}
     * @memberof TypeDate
     */
    year?: number;
    /**
     * Month of a year. Must be from 1 to 12, or 0 to specify a year without a
     * month and day.
     * @type {number}
     * @memberof TypeDate
     */
    month?: number;
    /**
     * Day of a month. Must be from 1 to 31 and valid for the year and month, or 0
     * to specify a year by itself or a year and month where the day isn't
     * significant.
     * @type {number}
     * @memberof TypeDate
     */
    day?: number;
}

export function TypeDateFromJSON(json: any): TypeDate {
    return TypeDateFromJSONTyped(json, false);
}

export function TypeDateFromJSONTyped(json: any, ignoreDiscriminator: boolean): TypeDate {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'year': !exists(json, 'year') ? undefined : json['year'],
        'month': !exists(json, 'month') ? undefined : json['month'],
        'day': !exists(json, 'day') ? undefined : json['day'],
    };
}

export function TypeDateToJSON(value?: TypeDate | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'year': value.year,
        'month': value.month,
        'day': value.day,
    };
}

