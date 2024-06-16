/**
 * Generated by orval v6.30.2 🍺
 * Do not edit manually.
 * h24s_10
 * traP 24 spring hackathon team 10
 * OpenAPI spec version: 0.0.1
 */
import {
  useMutation,
  useQuery
} from '@tanstack/vue-query'
import type {
  MutationFunction,
  QueryFunction,
  QueryKey,
  UseMutationOptions,
  UseMutationReturnType,
  UseQueryOptions,
  UseQueryReturnType
} from '@tanstack/vue-query'
import axios from 'axios'
import type {
  AxiosError,
  AxiosRequestConfig,
  AxiosResponse
} from 'axios'
import {
  computed,
  unref
} from 'vue'
import type {
  MaybeRef
} from 'vue'
export type GetEventsAllParams = {
includePastEvents?: boolean;
};

export interface TraQUser {
  displayName: string;
  name: string;
}

export interface TraQGroup {
  members?: TraQUser[];
  name?: string;
}

export interface Applicant {
  dateOptionIDs?: string[];
  traqID?: string;
}

export interface DateOption {
  end: string;
  id: string;
  start: string;
}

export interface GetAllEventsElement {
  end?: string;
  id: string;
  isConfirmed: boolean;
  start?: string;
  title: string;
}

export type GetTraQUsersResponse = TraQUser[];

export type GetTraQGroupsResponse = TraQGroup[];

export type GetEventApplicantsResponse = Applicant[];

export type GetEventTargetsResponse = string[];

export type GetEventParticipantsResponse = string[];

export interface PostEventApplicantsRequest {
  /** 何かコメントがあれば */
  comment: string;
  dateOptionIDs: string[];
}

export interface PatchEventConfirmRequest {
  eventDateOptionID?: string;
  isConfirmed: boolean;
}

export interface DateTimeResponse {
  end: string;
  start: string;
}

export interface GetEventResponse {
  date?: DateTimeResponse;
  dateOptions: DateOption[];
  description: string;
  /** traQ ID */
  hostID: string;
  id: string;
  isConfirmed: boolean;
  title: string;
}

export interface EventMeResponse {
  description?: string;
  event_id: string;
  isAnswered: boolean;
  isConfirmed: boolean;
  isHost: boolean;
  title: string;
}

export type EventMeResponses = EventMeResponse[];

export interface GetMyParticipateEvent {
  date: DateTimeResponse;
  description: string;
  id: string;
  title: string;
}

export type GetMyParticipateEventsResponse = GetMyParticipateEvent[];

export type GetAllEventsResponse = GetAllEventsElement[];

export interface PostEventResponse {
  id: string;
}

export type PostEventRequestDateOptionsItem = {
  end: string;
  start: string;
};

export interface PostEventRequest {
  dateOptions: PostEventRequestDateOptionsItem[];
  description: string;
  targets: string[];
  title: string;
}

export interface GetMeResponse {
  traQID?: string;
}





/**
 * Ping endpoint
 */
export const getPing = (
     options?: AxiosRequestConfig
 ): Promise<AxiosResponse<string>> => {
    
    return axios.get(
      `/api/ping`,options
    );
  }


export const getGetPingQueryKey = () => {
    return ['api','ping'] as const;
    }

    
export const getGetPingQueryOptions = <TData = Awaited<ReturnType<typeof getPing>>, TError = AxiosError<unknown>>( options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof getPing>>, TError, TData>>, axios?: AxiosRequestConfig}
) => {

const {query: queryOptions, axios: axiosOptions} = options ?? {};

  const queryKey =  getGetPingQueryKey();

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getPing>>> = ({ signal }) => getPing({ signal, ...axiosOptions });

      

      

   return  { queryKey, queryFn, ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof getPing>>, TError, TData> 
}

export type GetPingQueryResult = NonNullable<Awaited<ReturnType<typeof getPing>>>
export type GetPingQueryError = AxiosError<unknown>

export const useGetPing = <TData = Awaited<ReturnType<typeof getPing>>, TError = AxiosError<unknown>>(
  options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof getPing>>, TError, TData>>, axios?: AxiosRequestConfig}

  ): UseQueryReturnType<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetPingQueryOptions(options)

  const query = useQuery(queryOptions) as UseQueryReturnType<TData, TError> & { queryKey: QueryKey };

  query.queryKey = unref(queryOptions).queryKey as QueryKey;

  return query;
}




/**
 * get my traQ ID
 */
export const getMe = (
     options?: AxiosRequestConfig
 ): Promise<AxiosResponse<GetMeResponse>> => {
    
    return axios.get(
      `/api/me`,options
    );
  }


export const getGetMeQueryKey = () => {
    return ['api','me'] as const;
    }

    
export const getGetMeQueryOptions = <TData = Awaited<ReturnType<typeof getMe>>, TError = AxiosError<unknown>>( options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof getMe>>, TError, TData>>, axios?: AxiosRequestConfig}
) => {

const {query: queryOptions, axios: axiosOptions} = options ?? {};

  const queryKey =  getGetMeQueryKey();

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getMe>>> = ({ signal }) => getMe({ signal, ...axiosOptions });

      

      

   return  { queryKey, queryFn, ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof getMe>>, TError, TData> 
}

export type GetMeQueryResult = NonNullable<Awaited<ReturnType<typeof getMe>>>
export type GetMeQueryError = AxiosError<unknown>

export const useGetMe = <TData = Awaited<ReturnType<typeof getMe>>, TError = AxiosError<unknown>>(
  options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof getMe>>, TError, TData>>, axios?: AxiosRequestConfig}

  ): UseQueryReturnType<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetMeQueryOptions(options)

  const query = useQuery(queryOptions) as UseQueryReturnType<TData, TError> & { queryKey: QueryKey };

  query.queryKey = unref(queryOptions).queryKey as QueryKey;

  return query;
}




/**
 * create a new event
 */
export const postEvents = (
    postEventRequest: MaybeRef<PostEventRequest>, options?: AxiosRequestConfig
 ): Promise<AxiosResponse<PostEventResponse>> => {
    postEventRequest = unref(postEventRequest);
    return axios.post(
      `/api/events`,
      postEventRequest,options
    );
  }



export const getPostEventsMutationOptions = <TError = AxiosError<unknown>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof postEvents>>, TError,{data: PostEventRequest}, TContext>, axios?: AxiosRequestConfig}
): UseMutationOptions<Awaited<ReturnType<typeof postEvents>>, TError,{data: PostEventRequest}, TContext> => {
const {mutation: mutationOptions, axios: axiosOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof postEvents>>, {data: PostEventRequest}> = (props) => {
          const {data} = props ?? {};

          return  postEvents(data,axiosOptions)
        }

        


  return  { mutationFn, ...mutationOptions }}

    export type PostEventsMutationResult = NonNullable<Awaited<ReturnType<typeof postEvents>>>
    export type PostEventsMutationBody = PostEventRequest
    export type PostEventsMutationError = AxiosError<unknown>

    export const usePostEvents = <TError = AxiosError<unknown>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof postEvents>>, TError,{data: PostEventRequest}, TContext>, axios?: AxiosRequestConfig}
): UseMutationReturnType<
        Awaited<ReturnType<typeof postEvents>>,
        TError,
        {data: PostEventRequest},
        TContext
      > => {

      const mutationOptions = getPostEventsMutationOptions(options);

      return useMutation(mutationOptions);
    }
    
export const getEventsAll = (
    params?: MaybeRef<GetEventsAllParams>, options?: AxiosRequestConfig
 ): Promise<AxiosResponse<GetAllEventsResponse>> => {
    params = unref(params);
    return axios.get(
      `/api/events/all`,{
    ...options,
        params: {...unref(params), ...options?.params},}
    );
  }


export const getGetEventsAllQueryKey = (params?: MaybeRef<GetEventsAllParams>,) => {
    return ['api','events','all', ...(params ? [params]: [])] as const;
    }

    
export const getGetEventsAllQueryOptions = <TData = Awaited<ReturnType<typeof getEventsAll>>, TError = AxiosError<unknown>>(params?: MaybeRef<GetEventsAllParams>, options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof getEventsAll>>, TError, TData>>, axios?: AxiosRequestConfig}
) => {

const {query: queryOptions, axios: axiosOptions} = options ?? {};

  const queryKey =  getGetEventsAllQueryKey(params);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getEventsAll>>> = ({ signal }) => getEventsAll(params, { signal, ...axiosOptions });

      

      

   return  { queryKey, queryFn, ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof getEventsAll>>, TError, TData> 
}

export type GetEventsAllQueryResult = NonNullable<Awaited<ReturnType<typeof getEventsAll>>>
export type GetEventsAllQueryError = AxiosError<unknown>

export const useGetEventsAll = <TData = Awaited<ReturnType<typeof getEventsAll>>, TError = AxiosError<unknown>>(
 params?: MaybeRef<GetEventsAllParams>, options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof getEventsAll>>, TError, TData>>, axios?: AxiosRequestConfig}

  ): UseQueryReturnType<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetEventsAllQueryOptions(params,options)

  const query = useQuery(queryOptions) as UseQueryReturnType<TData, TError> & { queryKey: QueryKey };

  query.queryKey = unref(queryOptions).queryKey as QueryKey;

  return query;
}




export const getEventsMe = (
     options?: AxiosRequestConfig
 ): Promise<AxiosResponse<EventMeResponses>> => {
    
    return axios.get(
      `/api/events/me`,options
    );
  }


export const getGetEventsMeQueryKey = () => {
    return ['api','events','me'] as const;
    }

    
export const getGetEventsMeQueryOptions = <TData = Awaited<ReturnType<typeof getEventsMe>>, TError = AxiosError<unknown>>( options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof getEventsMe>>, TError, TData>>, axios?: AxiosRequestConfig}
) => {

const {query: queryOptions, axios: axiosOptions} = options ?? {};

  const queryKey =  getGetEventsMeQueryKey();

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getEventsMe>>> = ({ signal }) => getEventsMe({ signal, ...axiosOptions });

      

      

   return  { queryKey, queryFn, ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof getEventsMe>>, TError, TData> 
}

export type GetEventsMeQueryResult = NonNullable<Awaited<ReturnType<typeof getEventsMe>>>
export type GetEventsMeQueryError = AxiosError<unknown>

export const useGetEventsMe = <TData = Awaited<ReturnType<typeof getEventsMe>>, TError = AxiosError<unknown>>(
  options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof getEventsMe>>, TError, TData>>, axios?: AxiosRequestConfig}

  ): UseQueryReturnType<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetEventsMeQueryOptions(options)

  const query = useQuery(queryOptions) as UseQueryReturnType<TData, TError> & { queryKey: QueryKey };

  query.queryKey = unref(queryOptions).queryKey as QueryKey;

  return query;
}




/**
 * 自分の参加するイベント(isConfirmedがtrueのものだけ)
 */
export const getEventsMeParticipate = (
     options?: AxiosRequestConfig
 ): Promise<AxiosResponse<GetMyParticipateEventsResponse>> => {
    
    return axios.get(
      `/api/events/me/participate`,options
    );
  }


export const getGetEventsMeParticipateQueryKey = () => {
    return ['api','events','me','participate'] as const;
    }

    
export const getGetEventsMeParticipateQueryOptions = <TData = Awaited<ReturnType<typeof getEventsMeParticipate>>, TError = AxiosError<unknown>>( options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof getEventsMeParticipate>>, TError, TData>>, axios?: AxiosRequestConfig}
) => {

const {query: queryOptions, axios: axiosOptions} = options ?? {};

  const queryKey =  getGetEventsMeParticipateQueryKey();

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getEventsMeParticipate>>> = ({ signal }) => getEventsMeParticipate({ signal, ...axiosOptions });

      

      

   return  { queryKey, queryFn, ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof getEventsMeParticipate>>, TError, TData> 
}

export type GetEventsMeParticipateQueryResult = NonNullable<Awaited<ReturnType<typeof getEventsMeParticipate>>>
export type GetEventsMeParticipateQueryError = AxiosError<unknown>

export const useGetEventsMeParticipate = <TData = Awaited<ReturnType<typeof getEventsMeParticipate>>, TError = AxiosError<unknown>>(
  options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof getEventsMeParticipate>>, TError, TData>>, axios?: AxiosRequestConfig}

  ): UseQueryReturnType<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetEventsMeParticipateQueryOptions(options)

  const query = useQuery(queryOptions) as UseQueryReturnType<TData, TError> & { queryKey: QueryKey };

  query.queryKey = unref(queryOptions).queryKey as QueryKey;

  return query;
}




export const getEventsEventID = (
    eventID: MaybeRef<string>, options?: AxiosRequestConfig
 ): Promise<AxiosResponse<GetEventResponse>> => {
    eventID = unref(eventID);
    return axios.get(
      `/api/events/${eventID}`,options
    );
  }


export const getGetEventsEventIDQueryKey = (eventID: MaybeRef<string>,) => {
    return ['api','events',eventID] as const;
    }

    
export const getGetEventsEventIDQueryOptions = <TData = Awaited<ReturnType<typeof getEventsEventID>>, TError = AxiosError<unknown>>(eventID: MaybeRef<string>, options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof getEventsEventID>>, TError, TData>>, axios?: AxiosRequestConfig}
) => {

const {query: queryOptions, axios: axiosOptions} = options ?? {};

  const queryKey =  getGetEventsEventIDQueryKey(eventID);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getEventsEventID>>> = ({ signal }) => getEventsEventID(eventID, { signal, ...axiosOptions });

      

      

   return  { queryKey, queryFn, enabled: computed(() => !!(unref(eventID))), ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof getEventsEventID>>, TError, TData> 
}

export type GetEventsEventIDQueryResult = NonNullable<Awaited<ReturnType<typeof getEventsEventID>>>
export type GetEventsEventIDQueryError = AxiosError<unknown>

export const useGetEventsEventID = <TData = Awaited<ReturnType<typeof getEventsEventID>>, TError = AxiosError<unknown>>(
 eventID: MaybeRef<string>, options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof getEventsEventID>>, TError, TData>>, axios?: AxiosRequestConfig}

  ): UseQueryReturnType<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetEventsEventIDQueryOptions(eventID,options)

  const query = useQuery(queryOptions) as UseQueryReturnType<TData, TError> & { queryKey: QueryKey };

  query.queryKey = unref(queryOptions).queryKey as QueryKey;

  return query;
}




export const patchEventsEventIDConfirm = (
    eventID: MaybeRef<string>,
    patchEventConfirmRequest: MaybeRef<PatchEventConfirmRequest>, options?: AxiosRequestConfig
 ): Promise<AxiosResponse<void>> => {
    eventID = unref(eventID);
patchEventConfirmRequest = unref(patchEventConfirmRequest);
    return axios.patch(
      `/api/events/${eventID}/confirm`,
      patchEventConfirmRequest,options
    );
  }



export const getPatchEventsEventIDConfirmMutationOptions = <TError = AxiosError<unknown>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof patchEventsEventIDConfirm>>, TError,{eventID: string;data: PatchEventConfirmRequest}, TContext>, axios?: AxiosRequestConfig}
): UseMutationOptions<Awaited<ReturnType<typeof patchEventsEventIDConfirm>>, TError,{eventID: string;data: PatchEventConfirmRequest}, TContext> => {
const {mutation: mutationOptions, axios: axiosOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof patchEventsEventIDConfirm>>, {eventID: string;data: PatchEventConfirmRequest}> = (props) => {
          const {eventID,data} = props ?? {};

          return  patchEventsEventIDConfirm(eventID,data,axiosOptions)
        }

        


  return  { mutationFn, ...mutationOptions }}

    export type PatchEventsEventIDConfirmMutationResult = NonNullable<Awaited<ReturnType<typeof patchEventsEventIDConfirm>>>
    export type PatchEventsEventIDConfirmMutationBody = PatchEventConfirmRequest
    export type PatchEventsEventIDConfirmMutationError = AxiosError<unknown>

    export const usePatchEventsEventIDConfirm = <TError = AxiosError<unknown>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof patchEventsEventIDConfirm>>, TError,{eventID: string;data: PatchEventConfirmRequest}, TContext>, axios?: AxiosRequestConfig}
): UseMutationReturnType<
        Awaited<ReturnType<typeof patchEventsEventIDConfirm>>,
        TError,
        {eventID: string;data: PatchEventConfirmRequest},
        TContext
      > => {

      const mutationOptions = getPatchEventsEventIDConfirmMutationOptions(options);

      return useMutation(mutationOptions);
    }
    
/**
 * イベントの確定した参加者を取得
 */
export const getEventsEventIDParticipants = (
    eventID: MaybeRef<string>, options?: AxiosRequestConfig
 ): Promise<AxiosResponse<GetEventParticipantsResponse>> => {
    eventID = unref(eventID);
    return axios.get(
      `/api/events/${eventID}/participants`,options
    );
  }


export const getGetEventsEventIDParticipantsQueryKey = (eventID: MaybeRef<string>,) => {
    return ['api','events',eventID,'participants'] as const;
    }

    
export const getGetEventsEventIDParticipantsQueryOptions = <TData = Awaited<ReturnType<typeof getEventsEventIDParticipants>>, TError = AxiosError<unknown>>(eventID: MaybeRef<string>, options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof getEventsEventIDParticipants>>, TError, TData>>, axios?: AxiosRequestConfig}
) => {

const {query: queryOptions, axios: axiosOptions} = options ?? {};

  const queryKey =  getGetEventsEventIDParticipantsQueryKey(eventID);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getEventsEventIDParticipants>>> = ({ signal }) => getEventsEventIDParticipants(eventID, { signal, ...axiosOptions });

      

      

   return  { queryKey, queryFn, enabled: computed(() => !!(unref(eventID))), ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof getEventsEventIDParticipants>>, TError, TData> 
}

export type GetEventsEventIDParticipantsQueryResult = NonNullable<Awaited<ReturnType<typeof getEventsEventIDParticipants>>>
export type GetEventsEventIDParticipantsQueryError = AxiosError<unknown>

export const useGetEventsEventIDParticipants = <TData = Awaited<ReturnType<typeof getEventsEventIDParticipants>>, TError = AxiosError<unknown>>(
 eventID: MaybeRef<string>, options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof getEventsEventIDParticipants>>, TError, TData>>, axios?: AxiosRequestConfig}

  ): UseQueryReturnType<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetEventsEventIDParticipantsQueryOptions(eventID,options)

  const query = useQuery(queryOptions) as UseQueryReturnType<TData, TError> & { queryKey: QueryKey };

  query.queryKey = unref(queryOptions).queryKey as QueryKey;

  return query;
}




/**
 * イベントの参加候補者を取得
 */
export const getEventsEventIDTargets = (
    eventID: MaybeRef<string>, options?: AxiosRequestConfig
 ): Promise<AxiosResponse<GetEventTargetsResponse>> => {
    eventID = unref(eventID);
    return axios.get(
      `/api/events/${eventID}/targets`,options
    );
  }


export const getGetEventsEventIDTargetsQueryKey = (eventID: MaybeRef<string>,) => {
    return ['api','events',eventID,'targets'] as const;
    }

    
export const getGetEventsEventIDTargetsQueryOptions = <TData = Awaited<ReturnType<typeof getEventsEventIDTargets>>, TError = AxiosError<unknown>>(eventID: MaybeRef<string>, options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof getEventsEventIDTargets>>, TError, TData>>, axios?: AxiosRequestConfig}
) => {

const {query: queryOptions, axios: axiosOptions} = options ?? {};

  const queryKey =  getGetEventsEventIDTargetsQueryKey(eventID);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getEventsEventIDTargets>>> = ({ signal }) => getEventsEventIDTargets(eventID, { signal, ...axiosOptions });

      

      

   return  { queryKey, queryFn, enabled: computed(() => !!(unref(eventID))), ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof getEventsEventIDTargets>>, TError, TData> 
}

export type GetEventsEventIDTargetsQueryResult = NonNullable<Awaited<ReturnType<typeof getEventsEventIDTargets>>>
export type GetEventsEventIDTargetsQueryError = AxiosError<unknown>

export const useGetEventsEventIDTargets = <TData = Awaited<ReturnType<typeof getEventsEventIDTargets>>, TError = AxiosError<unknown>>(
 eventID: MaybeRef<string>, options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof getEventsEventIDTargets>>, TError, TData>>, axios?: AxiosRequestConfig}

  ): UseQueryReturnType<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetEventsEventIDTargetsQueryOptions(eventID,options)

  const query = useQuery(queryOptions) as UseQueryReturnType<TData, TError> & { queryKey: QueryKey };

  query.queryKey = unref(queryOptions).queryKey as QueryKey;

  return query;
}




/**
 * 都合のいい日程候補を提出
 */
export const postEventsEventIDApplicants = (
    eventID: MaybeRef<string>,
    postEventApplicantsRequest: MaybeRef<PostEventApplicantsRequest>, options?: AxiosRequestConfig
 ): Promise<AxiosResponse<void>> => {
    eventID = unref(eventID);
postEventApplicantsRequest = unref(postEventApplicantsRequest);
    return axios.post(
      `/api/events/${eventID}/applicants`,
      postEventApplicantsRequest,options
    );
  }



export const getPostEventsEventIDApplicantsMutationOptions = <TError = AxiosError<unknown>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof postEventsEventIDApplicants>>, TError,{eventID: string;data: PostEventApplicantsRequest}, TContext>, axios?: AxiosRequestConfig}
): UseMutationOptions<Awaited<ReturnType<typeof postEventsEventIDApplicants>>, TError,{eventID: string;data: PostEventApplicantsRequest}, TContext> => {
const {mutation: mutationOptions, axios: axiosOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof postEventsEventIDApplicants>>, {eventID: string;data: PostEventApplicantsRequest}> = (props) => {
          const {eventID,data} = props ?? {};

          return  postEventsEventIDApplicants(eventID,data,axiosOptions)
        }

        


  return  { mutationFn, ...mutationOptions }}

    export type PostEventsEventIDApplicantsMutationResult = NonNullable<Awaited<ReturnType<typeof postEventsEventIDApplicants>>>
    export type PostEventsEventIDApplicantsMutationBody = PostEventApplicantsRequest
    export type PostEventsEventIDApplicantsMutationError = AxiosError<unknown>

    export const usePostEventsEventIDApplicants = <TError = AxiosError<unknown>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof postEventsEventIDApplicants>>, TError,{eventID: string;data: PostEventApplicantsRequest}, TContext>, axios?: AxiosRequestConfig}
): UseMutationReturnType<
        Awaited<ReturnType<typeof postEventsEventIDApplicants>>,
        TError,
        {eventID: string;data: PostEventApplicantsRequest},
        TContext
      > => {

      const mutationOptions = getPostEventsEventIDApplicantsMutationOptions(options);

      return useMutation(mutationOptions);
    }
    
/**
 * イベント参加希望を出した人たちを取得
 */
export const getEventsEventIDApplicants = (
    eventID: MaybeRef<string>, options?: AxiosRequestConfig
 ): Promise<AxiosResponse<GetEventApplicantsResponse>> => {
    eventID = unref(eventID);
    return axios.get(
      `/api/events/${eventID}/applicants`,options
    );
  }


export const getGetEventsEventIDApplicantsQueryKey = (eventID: MaybeRef<string>,) => {
    return ['api','events',eventID,'applicants'] as const;
    }

    
export const getGetEventsEventIDApplicantsQueryOptions = <TData = Awaited<ReturnType<typeof getEventsEventIDApplicants>>, TError = AxiosError<unknown>>(eventID: MaybeRef<string>, options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof getEventsEventIDApplicants>>, TError, TData>>, axios?: AxiosRequestConfig}
) => {

const {query: queryOptions, axios: axiosOptions} = options ?? {};

  const queryKey =  getGetEventsEventIDApplicantsQueryKey(eventID);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getEventsEventIDApplicants>>> = ({ signal }) => getEventsEventIDApplicants(eventID, { signal, ...axiosOptions });

      

      

   return  { queryKey, queryFn, enabled: computed(() => !!(unref(eventID))), ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof getEventsEventIDApplicants>>, TError, TData> 
}

export type GetEventsEventIDApplicantsQueryResult = NonNullable<Awaited<ReturnType<typeof getEventsEventIDApplicants>>>
export type GetEventsEventIDApplicantsQueryError = AxiosError<unknown>

export const useGetEventsEventIDApplicants = <TData = Awaited<ReturnType<typeof getEventsEventIDApplicants>>, TError = AxiosError<unknown>>(
 eventID: MaybeRef<string>, options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof getEventsEventIDApplicants>>, TError, TData>>, axios?: AxiosRequestConfig}

  ): UseQueryReturnType<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetEventsEventIDApplicantsQueryOptions(eventID,options)

  const query = useQuery(queryOptions) as UseQueryReturnType<TData, TError> & { queryKey: QueryKey };

  query.queryKey = unref(queryOptions).queryKey as QueryKey;

  return query;
}




/**
 * get all traQ groups
 */
export const getTraqGroups = (
     options?: AxiosRequestConfig
 ): Promise<AxiosResponse<GetTraQGroupsResponse>> => {
    
    return axios.get(
      `/api/traq/groups`,options
    );
  }


export const getGetTraqGroupsQueryKey = () => {
    return ['api','traq','groups'] as const;
    }

    
export const getGetTraqGroupsQueryOptions = <TData = Awaited<ReturnType<typeof getTraqGroups>>, TError = AxiosError<unknown>>( options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof getTraqGroups>>, TError, TData>>, axios?: AxiosRequestConfig}
) => {

const {query: queryOptions, axios: axiosOptions} = options ?? {};

  const queryKey =  getGetTraqGroupsQueryKey();

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getTraqGroups>>> = ({ signal }) => getTraqGroups({ signal, ...axiosOptions });

      

      

   return  { queryKey, queryFn, ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof getTraqGroups>>, TError, TData> 
}

export type GetTraqGroupsQueryResult = NonNullable<Awaited<ReturnType<typeof getTraqGroups>>>
export type GetTraqGroupsQueryError = AxiosError<unknown>

export const useGetTraqGroups = <TData = Awaited<ReturnType<typeof getTraqGroups>>, TError = AxiosError<unknown>>(
  options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof getTraqGroups>>, TError, TData>>, axios?: AxiosRequestConfig}

  ): UseQueryReturnType<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetTraqGroupsQueryOptions(options)

  const query = useQuery(queryOptions) as UseQueryReturnType<TData, TError> & { queryKey: QueryKey };

  query.queryKey = unref(queryOptions).queryKey as QueryKey;

  return query;
}




/**
 * get all traQ users
 */
export const getTraqUsers = (
     options?: AxiosRequestConfig
 ): Promise<AxiosResponse<GetTraQUsersResponse>> => {
    
    return axios.get(
      `/api/traq/users`,options
    );
  }


export const getGetTraqUsersQueryKey = () => {
    return ['api','traq','users'] as const;
    }

    
export const getGetTraqUsersQueryOptions = <TData = Awaited<ReturnType<typeof getTraqUsers>>, TError = AxiosError<unknown>>( options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof getTraqUsers>>, TError, TData>>, axios?: AxiosRequestConfig}
) => {

const {query: queryOptions, axios: axiosOptions} = options ?? {};

  const queryKey =  getGetTraqUsersQueryKey();

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getTraqUsers>>> = ({ signal }) => getTraqUsers({ signal, ...axiosOptions });

      

      

   return  { queryKey, queryFn, ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof getTraqUsers>>, TError, TData> 
}

export type GetTraqUsersQueryResult = NonNullable<Awaited<ReturnType<typeof getTraqUsers>>>
export type GetTraqUsersQueryError = AxiosError<unknown>

export const useGetTraqUsers = <TData = Awaited<ReturnType<typeof getTraqUsers>>, TError = AxiosError<unknown>>(
  options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof getTraqUsers>>, TError, TData>>, axios?: AxiosRequestConfig}

  ): UseQueryReturnType<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetTraqUsersQueryOptions(options)

  const query = useQuery(queryOptions) as UseQueryReturnType<TData, TError> & { queryKey: QueryKey };

  query.queryKey = unref(queryOptions).queryKey as QueryKey;

  return query;
}




