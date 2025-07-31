import { API_URL } from "./chessApi";

let accessToken: string | null = null;
let refreshToken: string | null = null;
let refreshPromise: Promise<void> | null = null;

export const setAuthTokens = (access: string, refresh: string) => {
  accessToken = access;
  refreshToken = refresh;
};

export const clearAuthTokens = () => {
  accessToken = null;
  refreshToken = null;
};

async function refreshAuthToken(): Promise<void> {
  if (!refreshToken) {
    throw new Error("No refresh token");
  }

  try {
    const response = await fetch(`${API_URL}/api/refresh`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${refreshToken}`,
      },
    });

    if (!response.ok) {
      throw new Error("Error refreshing auth token");
    }

    const { access_token } = await response.json();
    accessToken = access_token;
  } catch (error) {
    clearAuthTokens();
    throw error;
  }
}

export async function authFetch(
  input: RequestInfo,
  init?: RequestInit
): Promise<Response> {
  const headers = new Headers(init?.headers);

  if (accessToken) {
    headers.set("Authorization", `Bearer ${accessToken}`);
  }

  let response = await fetch(input, {
    ...init,
    headers,
  });

  if (response.status === 401 && refreshToken) {
    if (!refreshPromise) {
      refreshPromise = refreshAuthToken().finally(() => {
        refreshPromise = null;
      });
    }

    await refreshPromise;

    if (accessToken) {
      headers.set("Authorization", `Bearer ${accessToken}`);
      response = await fetch(input, {
        ...init,
        headers,
      });
    }
  }

  return response;
}
