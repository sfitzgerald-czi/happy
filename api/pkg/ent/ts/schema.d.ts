/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */

export interface paths {
    "/app-configs": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        get: operations["listAppConfig"];
        put?: never;
        /** @description Sets an AppConfig with the specified Key/Value. */
        post: operations["setAppConfig"];
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/app-configs/{key}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** @description Finds the AppConfig with the requested Key and returns it. */
        get: operations["readAppConfig"];
        put?: never;
        post?: never;
        /** @description Deletes the AppConfig with the requested Key. */
        delete: operations["deleteAppConfig"];
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/health": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** Simple endpoint to check if the server is up */
        get: operations["Health"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
}
export type webhooks = Record<string, never>;
export interface components {
    schemas: {
        AppConfig: {
            /** Format: int64 */
            id: number;
            /** Format: date-time */
            created_at: string;
            /** Format: date-time */
            updated_at: string;
            /** Format: date-time */
            deleted_at?: string;
            app_name: string;
            environment: string;
            stack: string;
            key: string;
            value: string;
            /**
             * @default environment
             * @enum {string}
             */
            source: "stack" | "environment";
        };
        AppConfigList: {
            /** Format: int64 */
            id: number;
            /** Format: date-time */
            created_at: string;
            /** Format: date-time */
            updated_at: string;
            /** Format: date-time */
            deleted_at?: string;
            app_name: string;
            environment: string;
            stack: string;
            key: string;
            value: string;
            /**
             * @default environment
             * @enum {string}
             */
            source: "stack" | "environment";
        };
    };
    responses: {
        /** @description invalid input, data invalid */
        400: {
            headers: {
                [name: string]: unknown;
            };
            content: {
                "application/json": {
                    code: number;
                    status: string;
                    errors?: unknown;
                };
            };
        };
        /** @description insufficient permissions */
        403: {
            headers: {
                [name: string]: unknown;
            };
            content: {
                "application/json": {
                    code: number;
                    status: string;
                    errors?: unknown;
                };
            };
        };
        /** @description resource not found */
        404: {
            headers: {
                [name: string]: unknown;
            };
            content: {
                "application/json": {
                    code: number;
                    status: string;
                    errors?: unknown;
                };
            };
        };
        /** @description conflicting resources */
        409: {
            headers: {
                [name: string]: unknown;
            };
            content: {
                "application/json": {
                    code: number;
                    status: string;
                    errors?: unknown;
                };
            };
        };
        /** @description unexpected error */
        500: {
            headers: {
                [name: string]: unknown;
            };
            content: {
                "application/json": {
                    code: number;
                    status: string;
                    errors?: unknown;
                };
            };
        };
    };
    parameters: never;
    requestBodies: never;
    headers: never;
    pathItems: never;
}
export type $defs = Record<string, never>;
export interface operations {
    listAppConfig: {
        parameters: {
            query: {
                /** @description what page to render */
                page?: number;
                /** @description item count to render per page */
                itemsPerPage?: number;
                app_name: string;
                environment: string;
                stack?: string;
                aws_profile: string;
                aws_region: string;
                k8s_namespace: string;
                k8s_cluster_id: string;
            };
            header: {
                "X-Aws-Access-Key-Id": string;
                "X-Aws-Secret-Access-Key": string;
                "X-Aws-Session-Token": string;
            };
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description result AppConfig list */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["AppConfigList"][];
                };
            };
            400: components["responses"]["400"];
            403: components["responses"]["403"];
            404: components["responses"]["404"];
            409: components["responses"]["409"];
            500: components["responses"]["500"];
        };
    };
    setAppConfig: {
        parameters: {
            query: {
                /** @description what page to render */
                page?: number;
                /** @description item count to render per page */
                itemsPerPage?: number;
                app_name: string;
                environment: string;
                stack?: string;
                aws_profile: string;
                aws_region: string;
                k8s_namespace: string;
                k8s_cluster_id: string;
            };
            header: {
                "X-Aws-Access-Key-Id": string;
                "X-Aws-Secret-Access-Key": string;
                "X-Aws-Session-Token": string;
            };
            path?: never;
            cookie?: never;
        };
        /** @description AppConfig key/value to set */
        requestBody: {
            content: {
                "application/json": {
                    key: string;
                    value: string;
                };
            };
        };
        responses: {
            /** @description AppConfig with requested Key/Value was set */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["AppConfigList"];
                };
            };
            400: components["responses"]["400"];
            403: components["responses"]["403"];
            404: components["responses"]["404"];
            409: components["responses"]["409"];
            500: components["responses"]["500"];
        };
    };
    readAppConfig: {
        parameters: {
            query: {
                app_name: string;
                environment: string;
                stack?: string;
                aws_profile: string;
                aws_region: string;
                k8s_namespace: string;
                k8s_cluster_id: string;
            };
            header: {
                "X-Aws-Access-Key-Id": string;
                "X-Aws-Secret-Access-Key": string;
                "X-Aws-Session-Token": string;
            };
            path: {
                key: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description AppConfig with requested Key was found */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["AppConfigList"];
                };
            };
            400: components["responses"]["400"];
            403: components["responses"]["403"];
            404: components["responses"]["404"];
            409: components["responses"]["409"];
            500: components["responses"]["500"];
        };
    };
    deleteAppConfig: {
        parameters: {
            query: {
                app_name: string;
                environment: string;
                stack?: string;
                aws_profile: string;
                aws_region: string;
                k8s_namespace: string;
                k8s_cluster_id: string;
            };
            header: {
                "X-Aws-Access-Key-Id": string;
                "X-Aws-Secret-Access-Key": string;
                "X-Aws-Session-Token": string;
            };
            path: {
                key: string;
            };
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description AppConfig with requested Key was deleted */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
            400: components["responses"]["400"];
            403: components["responses"]["403"];
            404: components["responses"]["404"];
            409: components["responses"]["409"];
            500: components["responses"]["500"];
        };
    };
    Health: {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description Server is reachable */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": {
                        status: string;
                        route: string;
                        version: string;
                        git_sha: string;
                    };
                };
            };
            /** @description Server is not reachable */
            503: {
                headers: {
                    [name: string]: unknown;
                };
                content?: never;
            };
        };
    };
}
