export interface ErrorResponse {
  message: string;
}

export interface ValidationErrorResponse {
  validationErrors: string[];
}

export interface Department {
  id?: number;
  name: string;
  note?: string;
}
