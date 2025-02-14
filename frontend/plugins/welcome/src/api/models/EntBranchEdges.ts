/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntFaculty,
    EntFacultyFromJSON,
    EntFacultyFromJSONTyped,
    EntFacultyToJSON,
    EntUser,
    EntUserFromJSON,
    EntUserFromJSONTyped,
    EntUserToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntBranchEdges
 */
export interface EntBranchEdges {
    /**
     * 
     * @type {EntFaculty}
     * @memberof EntBranchEdges
     */
    faculty?: EntFaculty;
    /**
     * UserInformations holds the value of the user_informations edge.
     * @type {Array<EntUser>}
     * @memberof EntBranchEdges
     */
    userInformations?: Array<EntUser>;
}

export function EntBranchEdgesFromJSON(json: any): EntBranchEdges {
    return EntBranchEdgesFromJSONTyped(json, false);
}

export function EntBranchEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntBranchEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'faculty': !exists(json, 'faculty') ? undefined : EntFacultyFromJSON(json['faculty']),
        'userInformations': !exists(json, 'userInformations') ? undefined : ((json['userInformations'] as Array<any>).map(EntUserFromJSON)),
    };
}

export function EntBranchEdgesToJSON(value?: EntBranchEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'faculty': EntFacultyToJSON(value.faculty),
        'userInformations': value.userInformations === undefined ? undefined : ((value.userInformations as Array<any>).map(EntUserToJSON)),
    };
}


