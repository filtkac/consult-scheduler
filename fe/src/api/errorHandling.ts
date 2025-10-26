import { t } from '@lingui/core/macro';
import { ErrorResponse } from '@/api/model';

export async function handleFetchErrorResponse(response: Response) {
  const responseBody = await response.json();
  if (response.status === 400) {
    if (isErrorResponse(responseBody)) {
      throw new Error(responseBody.message);
    }
  }
  throw new Error(t`Server error. ${responseBody}`);
}

function isErrorResponse(body: unknown): body is ErrorResponse {
  return !!(body && typeof body === 'object' && 'messages' in body);
}
