export type User = {
  id: string;
  name: string;
  email: string;
  created_at: string;
  updated_at: string;
  access_token: string;
}

export type AuthContextType = {
  user: User | null;
  token: string | null;
  login: (email: string, password: string) => Promise<void>;
  register: (name: string, email: string, password: string) => Promise<void>;
  logout: () => void;
  isAuthenticated: boolean;
  loading: boolean;
}
