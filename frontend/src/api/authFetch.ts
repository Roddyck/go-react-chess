import { API_URL } from "./chessApi";

let accessToken = localStorage.getItem("accessToken");
let refreshToken = localStorage.getItem("refreshToken");
let refreshPromise: Promise<void> | null = null;

export const setAuthTokens = (access: string, refresh: string) => {
  localStorage.setItem("accessToken", access);
  localStorage.setItem("refreshToken", refresh);
};

export const clearAuthTokens = () => {
  localStorage.removeItem("accessToken");
  localStorage.removeItem("refreshToken");
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
    setAuthTokens(access_token, refreshToken);
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
  accessToken = localStorage.getItem("accessToken");
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
