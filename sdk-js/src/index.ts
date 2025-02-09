import * as Dto from "./dto";

const thorfinnApiPaths = {
  auth: {
    register: "/auth/register",
    verifyEmail: "/auth/verify-email",
    login: "/auth/login",
    logout: "/auth/logout",
    sendEmailVerification: "/auth/send-email-verification",
    sendPasswordResetLink: "/auth/send-password-reset-link",
    resetPassword: "/auth/reset-password",
    otpSend: "/auth/otp/send",
    otpVerify: "/auth/otp/verify",
  },
  user: {
    getAll: "/users",
    get: "/users/{id}",
    update: "/users/{id}",
    delete: "/users/{id}",
  },
};

export type Method = "GET" | "POST" | "PUT" | "DELETE";

export class ThorfinnError extends Error {
  status: number;
  details?: any;

  constructor(message: string, status: number, details?: any) {
    super(message);
    this.name = "ThorfinnError";
    this.status = status;
    this.details = details;
  }
}

class ThorfinnBaseClient {
  protected baseUrl: string;

  constructor(baseUrl: string) {
    this.baseUrl = baseUrl;
  }

  protected async makeRequest<T>(
    path: string,
    method: Method,
    body?: any
  ): Promise<T> {
    const response = await fetch(`${this.baseUrl}${path}`, {
      method: method,
      headers: {
        "Content-Type": "application/json",
      },
      body: body ? JSON.stringify(body) : undefined,
    });

    if (!response.ok) {
      const errorData = await response.json().catch(() => null);
      throw new ThorfinnError(
        errorData?.message || "Request failed",
        response.status,
        errorData
      );
    }

    return response.json() as Promise<T>;
  }
}

export class ThorfinnAuthClient extends ThorfinnBaseClient {
  async register(data: Dto.RegisterRequest): Promise<Dto.RegisterResponse> {
    return this.makeRequest<Dto.RegisterResponse>(
      thorfinnApiPaths.auth.register,
      "POST",
      data
    );
  }

  async verifyEmail(
    data: Dto.ConfirmEmailRequest
  ): Promise<Dto.ConfirmEmailResponse> {
    return this.makeRequest<Dto.ConfirmEmailResponse>(
      thorfinnApiPaths.auth.verifyEmail,
      "POST",
      data
    );
  }

  async login(data: Dto.LoginRequest): Promise<Dto.LoginResponse> {
    return this.makeRequest<Dto.LoginResponse>(
      thorfinnApiPaths.auth.login,
      "POST",
      data
    );
  }

  async logout(): Promise<Dto.LogoutResponse> {
    return this.makeRequest<Dto.LogoutResponse>(
      thorfinnApiPaths.auth.logout,
      "POST",
      {}
    );
  }

  async sendEmailVerification(
    data: Dto.SendVerificationEmailRequest
  ): Promise<Dto.SendVerificationEmailResponse> {
    return this.makeRequest<Dto.SendVerificationEmailResponse>(
      thorfinnApiPaths.auth.sendEmailVerification,
      "POST",
      data
    );
  }

  async sendPasswordResetLink(
    data: Dto.SendPasswordResetRequest
  ): Promise<Dto.SendPasswordResetResponse> {
    return this.makeRequest<Dto.SendPasswordResetResponse>(
      thorfinnApiPaths.auth.sendPasswordResetLink,
      "POST",
      data
    );
  }

  async resetPassword(
    data: Dto.ResetPasswordRequest
  ): Promise<Dto.ResetPasswordResponse> {
    return this.makeRequest<Dto.ResetPasswordResponse>(
      thorfinnApiPaths.auth.resetPassword,
      "POST",
      data
    );
  }

  async otpSend(data: Dto.OtpSendRequest): Promise<Dto.OtpSendResponse> {
    return this.makeRequest<Dto.OtpSendResponse>(
      thorfinnApiPaths.auth.otpSend,
      "POST",
      data
    );
  }

  async otpVerify(data: Dto.OtpVerifyRequest): Promise<Dto.OtpVerifyResponse> {
    return this.makeRequest<Dto.OtpVerifyResponse>(
      thorfinnApiPaths.auth.otpVerify,
      "POST",
      data
    );
  }
}

export class ThorfinnUserClient extends ThorfinnBaseClient {
  async getAllUsers(): Promise<Dto.GetAllUsersResponse> {
    return this.makeRequest<Dto.GetAllUsersResponse>(
      thorfinnApiPaths.user.getAll,
      "GET"
    );
  }

  async getUser(userId: string): Promise<Dto.GetUserResponse> {
    return this.makeRequest<Dto.GetUserResponse>(
      thorfinnApiPaths.user.get.replace("{id}", userId),
      "GET"
    );
  }

  async updateUser(
    userId: string,
    data: Dto.UpdateUserRequest
  ): Promise<Dto.UpdateUserResponse> {
    return this.makeRequest<Dto.UpdateUserResponse>(
      thorfinnApiPaths.user.update.replace("{id}", userId),
      "PUT",
      data
    );
  }

  async deleteUser(userId: string): Promise<Dto.DeleteUserResponse> {
    return this.makeRequest<Dto.DeleteUserResponse>(
      thorfinnApiPaths.user.delete.replace("{id}", userId),
      "DELETE"
    );
  }
}
