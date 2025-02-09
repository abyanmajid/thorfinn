export type RegisterRequest = {
  email: string;
  password: string;
  confirm_password: string;
};

export type RegisterResponse = {
  message: string;
};

export type ConfirmEmailRequest = {
  token: string;
};

export type ConfirmEmailResponse = {
  message: string;
};

export type LoginRequest = {
  email: string;
  password: string;
};

export type LoginResponse = {
  message: string;
  accessToken: string;
  refreshToken: string;
};

export type LogoutRequest = {};

export type LogoutResponse = {
  message: string;
};

export type SendVerificationEmailRequest = {
  email: string;
};

export type SendVerificationEmailResponse = {
  message: string;
};

export type SendPasswordResetRequest = {
  email: string;
};

export type SendPasswordResetResponse = {
  message: string;
};

export type ResetPasswordRequest = {
  token: string;
  newPassword: string;
};

export type ResetPasswordResponse = {
  message: string;
};

export type OtpSendRequest = {
  email: string;
};

export type OtpSendResponse = {
  message: string;
  otpCodeId: string;
};

export type OtpVerifyRequest = {
  otpCodeId: string;
  otpCode: string;
};

export type OtpVerifyResponse = {
  message: string;
};

export type ListUsersRow = {
  id: string;
  email: string;
  verified: boolean;
  twoFactorEnabled: boolean;
  createdAt: string;
  updatedAt: string;
};

export type ThorfinnUser = {
  id: string;
  email: string;
  passwordHash: string;
  verified: boolean;
  twoFactorEnabled: boolean;
  createdAt: string;
  updatedAt: string;
};

export type GetAllUsersRequest = {};

export type GetAllUsersResponse = {
  message: string;
  users: ListUsersRow[];
};

export type GetUserRequest = {};

export type GetUserResponse = {
  message: string;
  user: ThorfinnUser;
};

export type UpdateUserRequest = {
  email?: string;
  password?: string;
  verified?: boolean;
  twoFactorEnabled?: boolean;
};

export type UpdateUserResponse = {
  message: string;
};

export type DeleteUserRequest = {};

export type DeleteUserResponse = {
  message: string;
};
