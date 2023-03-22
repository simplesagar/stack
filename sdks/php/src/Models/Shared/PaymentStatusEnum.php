<?php

/**
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

declare(strict_types=1);

namespace formance\stack\Models\Shared;


enum PaymentStatusEnum: string
{
    case PENDING = 'PENDING';
    case ACTIVE = 'ACTIVE';
    case TERMINATED = 'TERMINATED';
    case FAILED = 'FAILED';
    case SUCCEEDED = 'SUCCEEDED';
    case CANCELLED = 'CANCELLED';
}
