// http types
type HttpMethod = "GET" | "POST" | "DELETE" | "PATCH";
type HttpHeaders = { [header in string]: string };
type HttpBody = string;
type HttpParams = URLSearchParams;
type HttpArgs = { [key in string]: any };
type FromHttpArgs = { body: HttpBody, params: HttpParams }

// Api options, result and type
class FetchApiArgs {
  endpoint: string;
  method?: HttpMethod;
  args?: HttpArgs;
};

class FetchApiResult {
  success: boolean;
  statusCode: number;
  message: string;
};

class FetchApi {
  fetch: (options: FetchApiArgs) => Promise<FetchApiResult>;
  getHeaders: () => HttpHeaders;
  setHeaders: (headers: HttpHeaders) => void;
}

// api factory 
const DefaultFromHttpArgs: (args: FetchApiArgs) => FromHttpArgs = (args) => {
  if (args.method == "GET") {
    console.log()
    return {
      body: "",
      params: new URLSearchParams({
        ...args.args
      })
    }
  }
  return {
    body: JSON.stringify(args.args ?? {}),
    params: new URLSearchParams()
  }
};

class ApiOptions {
  api: string;
  headers: HttpHeaders;
  retries?: 0 | 1 | 2 | 3;
  beforeRetry?: (options: ApiOptions, args: FetchApiArgs, result: FetchApiResult) => Promise<boolean>;
  fromArgs?: (args: FetchApiArgs) => FromHttpArgs;
}

export const NewApi = (options: ApiOptions): FetchApi => {
  options.retries ??= 0;
  options.fromArgs ??= DefaultFromHttpArgs;
  options.beforeRetry ??= async () => true;
  return {
    getHeaders: () => options.headers,
    setHeaders: (newHeaders) => (options.headers = newHeaders),
    fetch: async (args: FetchApiArgs): Promise<FetchApiResult> => {
      args.method ??= "GET";
      const {
        body,
        params
      } = options.fromArgs(args);

      let fetchesLeft = options.retries;
      let response: FetchApiResult;
      do {
        response = await fetch(`${options.api}/${args.endpoint}?${params.toString()}`, {
          method: args.method,
          headers: {
            "Content-Type": "application/json",
            ...options.headers
          },
          body: args.method == "GET" ? null : body,
        }).then(async res => ({
          success: res.ok,
          statusCode: res.status,
          message: await res.text(),
        }));
        if (
          response.success ||
          (fetchesLeft !== 0 && !(await options.beforeRetry(options, args, response)))
        ) return response;
      } while (fetchesLeft-- > 0);
      return response;
    },
  }
};