// @generated by protoc-gen-connect-es v0.12.0 with parameter "target=ts"
// @generated from file service/v1alpha1/service.proto (package akuity.io.kargo.service.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { AdminLoginRequest, AdminLoginResponse, CreateProjectRequest, CreateProjectResponse, CreatePromotionPolicyRequest, CreatePromotionPolicyResponse, CreateResourceRequest, CreateResourceResponse, CreateStageRequest, CreateStageResponse, DeleteProjectRequest, DeleteProjectResponse, DeletePromotionPolicyRequest, DeletePromotionPolicyResponse, DeleteResourceRequest, DeleteResourceResponse, DeleteStageRequest, DeleteStageResponse, GetPromotionPolicyRequest, GetPromotionPolicyResponse, GetPromotionRequest, GetPromotionResponse, GetPublicConfigRequest, GetPublicConfigResponse, GetStageRequest, GetStageResponse, GetVersionInfoRequest, GetVersionInfoResponse, ListProjectsRequest, ListProjectsResponse, ListPromotionPoliciesRequest, ListPromotionPoliciesResponse, ListPromotionsRequest, ListPromotionsResponse, ListStagesRequest, ListStagesResponse, PromoteStageRequest, PromoteStageResponse, PromoteSubscribersRequest, PromoteSubscribersResponse, QueryFreightRequest, QueryFreightResponse, RefreshStageRequest, RefreshStageResponse, SetAutoPromotionForStageRequest, SetAutoPromotionForStageResponse, UpdatePromotionPolicyRequest, UpdatePromotionPolicyResponse, UpdateResourceRequest, UpdateResourceResponse, UpdateStageRequest, UpdateStageResponse, WatchStagesRequest, WatchStagesResponse } from "./service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service akuity.io.kargo.service.v1alpha1.KargoService
 */
export const KargoService = {
  typeName: "akuity.io.kargo.service.v1alpha1.KargoService",
  methods: {
    /**
     * @generated from rpc akuity.io.kargo.service.v1alpha1.KargoService.GetVersionInfo
     */
    getVersionInfo: {
      name: "GetVersionInfo",
      I: GetVersionInfoRequest,
      O: GetVersionInfoResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc akuity.io.kargo.service.v1alpha1.KargoService.GetPublicConfig
     */
    getPublicConfig: {
      name: "GetPublicConfig",
      I: GetPublicConfigRequest,
      O: GetPublicConfigResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc akuity.io.kargo.service.v1alpha1.KargoService.AdminLogin
     */
    adminLogin: {
      name: "AdminLogin",
      I: AdminLoginRequest,
      O: AdminLoginResponse,
      kind: MethodKind.Unary,
    },
    /**
     * TODO(devholic): Add ApplyResource API
     * rpc ApplyResource(ApplyResourceRequest) returns (ApplyResourceRequest);
     *
     * @generated from rpc akuity.io.kargo.service.v1alpha1.KargoService.CreateResource
     */
    createResource: {
      name: "CreateResource",
      I: CreateResourceRequest,
      O: CreateResourceResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc akuity.io.kargo.service.v1alpha1.KargoService.UpdateResource
     */
    updateResource: {
      name: "UpdateResource",
      I: UpdateResourceRequest,
      O: UpdateResourceResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc akuity.io.kargo.service.v1alpha1.KargoService.DeleteResource
     */
    deleteResource: {
      name: "DeleteResource",
      I: DeleteResourceRequest,
      O: DeleteResourceResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc akuity.io.kargo.service.v1alpha1.KargoService.CreateStage
     */
    createStage: {
      name: "CreateStage",
      I: CreateStageRequest,
      O: CreateStageResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc akuity.io.kargo.service.v1alpha1.KargoService.ListStages
     */
    listStages: {
      name: "ListStages",
      I: ListStagesRequest,
      O: ListStagesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc akuity.io.kargo.service.v1alpha1.KargoService.GetStage
     */
    getStage: {
      name: "GetStage",
      I: GetStageRequest,
      O: GetStageResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc akuity.io.kargo.service.v1alpha1.KargoService.WatchStages
     */
    watchStages: {
      name: "WatchStages",
      I: WatchStagesRequest,
      O: WatchStagesResponse,
      kind: MethodKind.ServerStreaming,
    },
    /**
     * @generated from rpc akuity.io.kargo.service.v1alpha1.KargoService.UpdateStage
     */
    updateStage: {
      name: "UpdateStage",
      I: UpdateStageRequest,
      O: UpdateStageResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc akuity.io.kargo.service.v1alpha1.KargoService.DeleteStage
     */
    deleteStage: {
      name: "DeleteStage",
      I: DeleteStageRequest,
      O: DeleteStageResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc akuity.io.kargo.service.v1alpha1.KargoService.PromoteStage
     */
    promoteStage: {
      name: "PromoteStage",
      I: PromoteStageRequest,
      O: PromoteStageResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc akuity.io.kargo.service.v1alpha1.KargoService.PromoteSubscribers
     */
    promoteSubscribers: {
      name: "PromoteSubscribers",
      I: PromoteSubscribersRequest,
      O: PromoteSubscribersResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc akuity.io.kargo.service.v1alpha1.KargoService.RefreshStage
     */
    refreshStage: {
      name: "RefreshStage",
      I: RefreshStageRequest,
      O: RefreshStageResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Promotion APIs 
     *
     * @generated from rpc akuity.io.kargo.service.v1alpha1.KargoService.ListPromotions
     */
    listPromotions: {
      name: "ListPromotions",
      I: ListPromotionsRequest,
      O: ListPromotionsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc akuity.io.kargo.service.v1alpha1.KargoService.GetPromotion
     */
    getPromotion: {
      name: "GetPromotion",
      I: GetPromotionRequest,
      O: GetPromotionResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc akuity.io.kargo.service.v1alpha1.KargoService.SetAutoPromotionForStage
     */
    setAutoPromotionForStage: {
      name: "SetAutoPromotionForStage",
      I: SetAutoPromotionForStageRequest,
      O: SetAutoPromotionForStageResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc akuity.io.kargo.service.v1alpha1.KargoService.CreatePromotionPolicy
     */
    createPromotionPolicy: {
      name: "CreatePromotionPolicy",
      I: CreatePromotionPolicyRequest,
      O: CreatePromotionPolicyResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc akuity.io.kargo.service.v1alpha1.KargoService.ListPromotionPolicies
     */
    listPromotionPolicies: {
      name: "ListPromotionPolicies",
      I: ListPromotionPoliciesRequest,
      O: ListPromotionPoliciesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc akuity.io.kargo.service.v1alpha1.KargoService.GetPromotionPolicy
     */
    getPromotionPolicy: {
      name: "GetPromotionPolicy",
      I: GetPromotionPolicyRequest,
      O: GetPromotionPolicyResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc akuity.io.kargo.service.v1alpha1.KargoService.UpdatePromotionPolicy
     */
    updatePromotionPolicy: {
      name: "UpdatePromotionPolicy",
      I: UpdatePromotionPolicyRequest,
      O: UpdatePromotionPolicyResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc akuity.io.kargo.service.v1alpha1.KargoService.DeletePromotionPolicy
     */
    deletePromotionPolicy: {
      name: "DeletePromotionPolicy",
      I: DeletePromotionPolicyRequest,
      O: DeletePromotionPolicyResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc akuity.io.kargo.service.v1alpha1.KargoService.CreateProject
     */
    createProject: {
      name: "CreateProject",
      I: CreateProjectRequest,
      O: CreateProjectResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc akuity.io.kargo.service.v1alpha1.KargoService.ListProjects
     */
    listProjects: {
      name: "ListProjects",
      I: ListProjectsRequest,
      O: ListProjectsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc akuity.io.kargo.service.v1alpha1.KargoService.DeleteProject
     */
    deleteProject: {
      name: "DeleteProject",
      I: DeleteProjectRequest,
      O: DeleteProjectResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc akuity.io.kargo.service.v1alpha1.KargoService.QueryFreight
     */
    queryFreight: {
      name: "QueryFreight",
      I: QueryFreightRequest,
      O: QueryFreightResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

