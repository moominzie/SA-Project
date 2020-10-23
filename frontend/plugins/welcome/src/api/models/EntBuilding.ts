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
    EntBuildingEdges,
    EntBuildingEdgesFromJSON,
    EntBuildingEdgesFromJSONTyped,
    EntBuildingEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntBuilding
 */
export interface EntBuilding {
    /**
     * Buname holds the value of the "buname" field.
     * @type {string}
     * @memberof EntBuilding
     */
    buname?: string;
    /**
     * 
     * @type {EntBuildingEdges}
     * @memberof EntBuilding
     */
    edges?: EntBuildingEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntBuilding
     */
    id?: number;
}

export function EntBuildingFromJSON(json: any): EntBuilding {
    return EntBuildingFromJSONTyped(json, false);
}

export function EntBuildingFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntBuilding {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'buname': !exists(json, 'buname') ? undefined : json['buname'],
        'edges': !exists(json, 'edges') ? undefined : EntBuildingEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntBuildingToJSON(value?: EntBuilding | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'buname': value.buname,
        'edges': EntBuildingEdgesToJSON(value.edges),
        'id': value.id,
    };
}


