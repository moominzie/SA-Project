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


import * as runtime from '../runtime';
import {
    ControllersUser,
    ControllersUserFromJSON,
    ControllersUserToJSON,
    EntBranch,
    EntBranchFromJSON,
    EntBranchToJSON,
    EntBuilding,
    EntBuildingFromJSON,
    EntBuildingToJSON,
    EntEmployee,
    EntEmployeeFromJSON,
    EntEmployeeToJSON,
    EntFaculty,
    EntFacultyFromJSON,
    EntFacultyToJSON,
    EntRoom,
    EntRoomFromJSON,
    EntRoomToJSON,
    EntUser,
    EntUserFromJSON,
    EntUserToJSON,
} from '../models';

export interface CreateBranchRequest {
    branch: EntBranch;
}

export interface CreateBuildingRequest {
    building: EntBuilding;
}

export interface CreateEmployeeRequest {
    employee: EntEmployee;
}

export interface CreateFacultyRequest {
    faculty: EntFaculty;
}

export interface CreateRoomRequest {
    room: EntRoom;
}

export interface CreateUserRequest {
    user: ControllersUser;
}

export interface DeleteBranchRequest {
    id: number;
}

export interface DeleteBuildingRequest {
    id: number;
}

export interface DeleteEmployeeRequest {
    id: number;
}

export interface DeleteFacultyRequest {
    id: number;
}

export interface DeleteRoomRequest {
    id: number;
}

export interface DeleteUserRequest {
    id: number;
}

export interface GetBranchRequest {
    id: number;
}

export interface GetBuildingRequest {
    id: number;
}

export interface GetEmployeeRequest {
    id: number;
}

export interface GetFacultyRequest {
    id: number;
}

export interface GetRoomRequest {
    id: number;
}

export interface ListBranchRequest {
    limit?: number;
    offset?: number;
}

export interface ListBuildingRequest {
    limit?: number;
    offset?: number;
}

export interface ListEmployeeRequest {
    limit?: number;
    offset?: number;
}

export interface ListFacultyRequest {
    limit?: number;
    offset?: number;
}

export interface ListRoomRequest {
    limit?: number;
    offset?: number;
}

export interface ListUserRequest {
    limit?: number;
    offset?: number;
}

/**
 * 
 */
export class DefaultApi extends runtime.BaseAPI {

    /**
     * Create branch
     * Create branch
     */
    async createBranchRaw(requestParameters: CreateBranchRequest): Promise<runtime.ApiResponse<EntBranch>> {
        if (requestParameters.branch === null || requestParameters.branch === undefined) {
            throw new runtime.RequiredError('branch','Required parameter requestParameters.branch was null or undefined when calling createBranch.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/branchs`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntBranchToJSON(requestParameters.branch),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntBranchFromJSON(jsonValue));
    }

    /**
     * Create branch
     * Create branch
     */
    async createBranch(requestParameters: CreateBranchRequest): Promise<EntBranch> {
        const response = await this.createBranchRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create building
     * Create building
     */
    async createBuildingRaw(requestParameters: CreateBuildingRequest): Promise<runtime.ApiResponse<EntBuilding>> {
        if (requestParameters.building === null || requestParameters.building === undefined) {
            throw new runtime.RequiredError('building','Required parameter requestParameters.building was null or undefined when calling createBuilding.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/buildings`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntBuildingToJSON(requestParameters.building),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntBuildingFromJSON(jsonValue));
    }

    /**
     * Create building
     * Create building
     */
    async createBuilding(requestParameters: CreateBuildingRequest): Promise<EntBuilding> {
        const response = await this.createBuildingRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create employee
     * Create employee
     */
    async createEmployeeRaw(requestParameters: CreateEmployeeRequest): Promise<runtime.ApiResponse<EntEmployee>> {
        if (requestParameters.employee === null || requestParameters.employee === undefined) {
            throw new runtime.RequiredError('employee','Required parameter requestParameters.employee was null or undefined when calling createEmployee.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/employees`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntEmployeeToJSON(requestParameters.employee),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntEmployeeFromJSON(jsonValue));
    }

    /**
     * Create employee
     * Create employee
     */
    async createEmployee(requestParameters: CreateEmployeeRequest): Promise<EntEmployee> {
        const response = await this.createEmployeeRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create faculty
     * Create faculty
     */
    async createFacultyRaw(requestParameters: CreateFacultyRequest): Promise<runtime.ApiResponse<EntFaculty>> {
        if (requestParameters.faculty === null || requestParameters.faculty === undefined) {
            throw new runtime.RequiredError('faculty','Required parameter requestParameters.faculty was null or undefined when calling createFaculty.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/facultys`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntFacultyToJSON(requestParameters.faculty),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntFacultyFromJSON(jsonValue));
    }

    /**
     * Create faculty
     * Create faculty
     */
    async createFaculty(requestParameters: CreateFacultyRequest): Promise<EntFaculty> {
        const response = await this.createFacultyRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create room
     * Create room
     */
    async createRoomRaw(requestParameters: CreateRoomRequest): Promise<runtime.ApiResponse<EntRoom>> {
        if (requestParameters.room === null || requestParameters.room === undefined) {
            throw new runtime.RequiredError('room','Required parameter requestParameters.room was null or undefined when calling createRoom.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/rooms`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntRoomToJSON(requestParameters.room),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntRoomFromJSON(jsonValue));
    }

    /**
     * Create room
     * Create room
     */
    async createRoom(requestParameters: CreateRoomRequest): Promise<EntRoom> {
        const response = await this.createRoomRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create user
     * Create user
     */
    async createUserRaw(requestParameters: CreateUserRequest): Promise<runtime.ApiResponse<ControllersUser>> {
        if (requestParameters.user === null || requestParameters.user === undefined) {
            throw new runtime.RequiredError('user','Required parameter requestParameters.user was null or undefined when calling createUser.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/users`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ControllersUserToJSON(requestParameters.user),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => ControllersUserFromJSON(jsonValue));
    }

    /**
     * Create user
     * Create user
     */
    async createUser(requestParameters: CreateUserRequest): Promise<ControllersUser> {
        const response = await this.createUserRaw(requestParameters);
        return await response.value();
    }

    /**
     * get branch by ID
     * Delete a branch entity by ID
     */
    async deleteBranchRaw(requestParameters: DeleteBranchRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deleteBranch.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/branch/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * get branch by ID
     * Delete a branch entity by ID
     */
    async deleteBranch(requestParameters: DeleteBranchRequest): Promise<object> {
        const response = await this.deleteBranchRaw(requestParameters);
        return await response.value();
    }

    /**
     * get building by ID
     * Delete a building entity by ID
     */
    async deleteBuildingRaw(requestParameters: DeleteBuildingRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deleteBuilding.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/building/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * get building by ID
     * Delete a building entity by ID
     */
    async deleteBuilding(requestParameters: DeleteBuildingRequest): Promise<object> {
        const response = await this.deleteBuildingRaw(requestParameters);
        return await response.value();
    }

    /**
     * get employee by ID
     * Delete a employee entity by ID
     */
    async deleteEmployeeRaw(requestParameters: DeleteEmployeeRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deleteEmployee.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/employee/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * get employee by ID
     * Delete a employee entity by ID
     */
    async deleteEmployee(requestParameters: DeleteEmployeeRequest): Promise<object> {
        const response = await this.deleteEmployeeRaw(requestParameters);
        return await response.value();
    }

    /**
     * get faculty by ID
     * Delete a faculty entity by ID
     */
    async deleteFacultyRaw(requestParameters: DeleteFacultyRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deleteFaculty.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/faculty/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * get faculty by ID
     * Delete a faculty entity by ID
     */
    async deleteFaculty(requestParameters: DeleteFacultyRequest): Promise<object> {
        const response = await this.deleteFacultyRaw(requestParameters);
        return await response.value();
    }

    /**
     * get room by ID
     * Delete a room entity by ID
     */
    async deleteRoomRaw(requestParameters: DeleteRoomRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deleteRoom.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/room/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * get room by ID
     * Delete a room entity by ID
     */
    async deleteRoom(requestParameters: DeleteRoomRequest): Promise<object> {
        const response = await this.deleteRoomRaw(requestParameters);
        return await response.value();
    }

    /**
     * get user by ID
     * Delete a user entity by ID
     */
    async deleteUserRaw(requestParameters: DeleteUserRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deleteUser.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/users/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * get user by ID
     * Delete a user entity by ID
     */
    async deleteUser(requestParameters: DeleteUserRequest): Promise<object> {
        const response = await this.deleteUserRaw(requestParameters);
        return await response.value();
    }

    /**
     * get branch by ID
     * Get a branch entity by ID
     */
    async getBranchRaw(requestParameters: GetBranchRequest): Promise<runtime.ApiResponse<EntBranch>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getBranch.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/branchs/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntBranchFromJSON(jsonValue));
    }

    /**
     * get branch by ID
     * Get a branch entity by ID
     */
    async getBranch(requestParameters: GetBranchRequest): Promise<EntBranch> {
        const response = await this.getBranchRaw(requestParameters);
        return await response.value();
    }

    /**
     * get building by ID
     * Get a building entity by ID
     */
    async getBuildingRaw(requestParameters: GetBuildingRequest): Promise<runtime.ApiResponse<EntBuilding>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getBuilding.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/buildings/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntBuildingFromJSON(jsonValue));
    }

    /**
     * get building by ID
     * Get a building entity by ID
     */
    async getBuilding(requestParameters: GetBuildingRequest): Promise<EntBuilding> {
        const response = await this.getBuildingRaw(requestParameters);
        return await response.value();
    }

    /**
     * get employee by ID
     * Get a employee entity by ID
     */
    async getEmployeeRaw(requestParameters: GetEmployeeRequest): Promise<runtime.ApiResponse<EntEmployee>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getEmployee.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/employees/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntEmployeeFromJSON(jsonValue));
    }

    /**
     * get employee by ID
     * Get a employee entity by ID
     */
    async getEmployee(requestParameters: GetEmployeeRequest): Promise<EntEmployee> {
        const response = await this.getEmployeeRaw(requestParameters);
        return await response.value();
    }

    /**
     * get faculty by ID
     * Get a faculty entity by ID
     */
    async getFacultyRaw(requestParameters: GetFacultyRequest): Promise<runtime.ApiResponse<EntFaculty>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getFaculty.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/facultys/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntFacultyFromJSON(jsonValue));
    }

    /**
     * get faculty by ID
     * Get a faculty entity by ID
     */
    async getFaculty(requestParameters: GetFacultyRequest): Promise<EntFaculty> {
        const response = await this.getFacultyRaw(requestParameters);
        return await response.value();
    }

    /**
     * get room by ID
     * Get a room entity by ID
     */
    async getRoomRaw(requestParameters: GetRoomRequest): Promise<runtime.ApiResponse<EntRoom>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getRoom.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/rooms/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntRoomFromJSON(jsonValue));
    }

    /**
     * get room by ID
     * Get a room entity by ID
     */
    async getRoom(requestParameters: GetRoomRequest): Promise<EntRoom> {
        const response = await this.getRoomRaw(requestParameters);
        return await response.value();
    }

    /**
     * list branch entities
     * List branch entities
     */
    async listBranchRaw(requestParameters: ListBranchRequest): Promise<runtime.ApiResponse<Array<EntBranch>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/branchs`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntBranchFromJSON));
    }

    /**
     * list branch entities
     * List branch entities
     */
    async listBranch(requestParameters: ListBranchRequest): Promise<Array<EntBranch>> {
        const response = await this.listBranchRaw(requestParameters);
        return await response.value();
    }

    /**
     * list building entities
     * List building entities
     */
    async listBuildingRaw(requestParameters: ListBuildingRequest): Promise<runtime.ApiResponse<Array<EntBuilding>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/buildings`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntBuildingFromJSON));
    }

    /**
     * list building entities
     * List building entities
     */
    async listBuilding(requestParameters: ListBuildingRequest): Promise<Array<EntBuilding>> {
        const response = await this.listBuildingRaw(requestParameters);
        return await response.value();
    }

    /**
     * list employee entities
     * List employee entities
     */
    async listEmployeeRaw(requestParameters: ListEmployeeRequest): Promise<runtime.ApiResponse<Array<EntEmployee>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/employees`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntEmployeeFromJSON));
    }

    /**
     * list employee entities
     * List employee entities
     */
    async listEmployee(requestParameters: ListEmployeeRequest): Promise<Array<EntEmployee>> {
        const response = await this.listEmployeeRaw(requestParameters);
        return await response.value();
    }

    /**
     * list faculty entities
     * List faculty entities
     */
    async listFacultyRaw(requestParameters: ListFacultyRequest): Promise<runtime.ApiResponse<Array<EntFaculty>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/facultys`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntFacultyFromJSON));
    }

    /**
     * list faculty entities
     * List faculty entities
     */
    async listFaculty(requestParameters: ListFacultyRequest): Promise<Array<EntFaculty>> {
        const response = await this.listFacultyRaw(requestParameters);
        return await response.value();
    }

    /**
     * list room entities
     * List room entities
     */
    async listRoomRaw(requestParameters: ListRoomRequest): Promise<runtime.ApiResponse<Array<EntRoom>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/rooms`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntRoomFromJSON));
    }

    /**
     * list room entities
     * List room entities
     */
    async listRoom(requestParameters: ListRoomRequest): Promise<Array<EntRoom>> {
        const response = await this.listRoomRaw(requestParameters);
        return await response.value();
    }

    /**
     * list user entities
     * List user entities
     */
    async listUserRaw(requestParameters: ListUserRequest): Promise<runtime.ApiResponse<Array<EntUser>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/users`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntUserFromJSON));
    }

    /**
     * list user entities
     * List user entities
     */
    async listUser(requestParameters: ListUserRequest): Promise<Array<EntUser>> {
        const response = await this.listUserRaw(requestParameters);
        return await response.value();
    }

}
