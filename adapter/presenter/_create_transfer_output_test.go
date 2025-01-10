// ********RoostGPT********
/*
Test generated by RoostGPT for test mayankzbio-go-clean using AI Type Claude AI and AI Model claude-3-5-sonnet-20240620

ROOST_METHOD_HASH=Output_b4dd5ccd4f
ROOST_METHOD_SIG_HASH=Output_639eeea89e

FUNCTION_DEF=func (c createTransferPresenter) Output(transfer domain.Transfer) usecase.CreateTransferOutput
Based on the provided function and context, here are several test scenarios for the `Output` method of the `createTransferPresenter` struct:

```
Scenario 1: Valid Transfer Output

Details:
  Description: Test that the Output method correctly transforms a valid domain.Transfer object into a usecase.CreateTransferOutput object with all fields properly mapped.

Execution:
  Arrange:
    - Create a mock domain.Transfer object with known values for all fields.
    - Initialize a createTransferPresenter instance.
  Act:
    - Call the Output method with the mock Transfer object.
  Assert:
    - Verify that the returned CreateTransferOutput object has all fields matching the input Transfer object.
    - Check that the CreatedAt field is correctly formatted as RFC3339.

Validation:
  This test ensures that the basic functionality of the Output method works as expected, correctly mapping all fields from the domain model to the use case output model. It's crucial for maintaining the consistency between the internal domain representation and the external API response.

Scenario 2: Zero Amount Transfer

Details:
  Description: Verify that the Output method correctly handles a Transfer with a zero amount.

Execution:
  Arrange:
    - Create a mock domain.Transfer object with an amount of 0.
    - Initialize a createTransferPresenter instance.
  Act:
    - Call the Output method with the zero-amount Transfer object.
  Assert:
    - Verify that the Amount field in the returned CreateTransferOutput is exactly 0.0.

Validation:
  This test checks the edge case of a zero-amount transfer, ensuring that the presenter doesn't modify or reject such transfers. It's important for scenarios where zero-amount transfers might be valid in the business logic.

Scenario 3: Large Amount Transfer

Details:
  Description: Test the Output method's handling of a Transfer with a very large amount to check for potential floating-point precision issues.

Execution:
  Arrange:
    - Create a mock domain.Transfer object with a very large amount (e.g., 999999999.99).
    - Initialize a createTransferPresenter instance.
  Act:
    - Call the Output method with the large-amount Transfer object.
  Assert:
    - Verify that the Amount field in the returned CreateTransferOutput matches the input amount exactly.

Validation:
  This test ensures that the presenter can handle large monetary values without loss of precision. It's crucial for financial applications where accuracy is paramount.

Scenario 4: Transfer with Minimum Valid IDs

Details:
  Description: Verify that the Output method correctly handles a Transfer with minimum length valid IDs for all ID fields.

Execution:
  Arrange:
    - Create a mock domain.Transfer object with minimum length valid IDs for TransferID, AccountOriginID, and AccountDestinationID.
    - Initialize a createTransferPresenter instance.
  Act:
    - Call the Output method with the prepared Transfer object.
  Assert:
    - Verify that all ID fields in the returned CreateTransferOutput are correctly stringified versions of the input IDs.

Validation:
  This test checks the presenter's ability to handle edge cases with minimal valid IDs, ensuring that no truncation or modification occurs during the transformation process.

Scenario 5: Transfer with Maximum Valid IDs

Details:
  Description: Test the Output method's handling of a Transfer with maximum length valid IDs for all ID fields.

Execution:
  Arrange:
    - Create a mock domain.Transfer object with maximum length valid IDs for TransferID, AccountOriginID, and AccountDestinationID.
    - Initialize a createTransferPresenter instance.
  Act:
    - Call the Output method with the prepared Transfer object.
  Assert:
    - Verify that all ID fields in the returned CreateTransferOutput are correctly stringified versions of the input IDs, without any truncation.

Validation:
  This test ensures that the presenter can handle large IDs without issues, which is important for maintaining data integrity and avoiding potential bugs with ID handling in the system.

Scenario 6: Transfer with Past Timestamp

Details:
  Description: Verify that the Output method correctly formats a Transfer with a past timestamp.

Execution:
  Arrange:
    - Create a mock domain.Transfer object with a CreatedAt time set to a known past date and time.
    - Initialize a createTransferPresenter instance.
  Act:
    - Call the Output method with the prepared Transfer object.
  Assert:
    - Verify that the CreatedAt field in the returned CreateTransferOutput is correctly formatted as RFC3339 and matches the input time.

Validation:
  This test ensures that the presenter correctly handles and formats past timestamps, which is crucial for maintaining accurate historical records of transfers.

Scenario 7: Transfer with Future Timestamp

Details:
  Description: Test the Output method's handling of a Transfer with a future timestamp.

Execution:
  Arrange:
    - Create a mock domain.Transfer object with a CreatedAt time set to a known future date and time.
    - Initialize a createTransferPresenter instance.
  Act:
    - Call the Output method with the prepared Transfer object.
  Assert:
    - Verify that the CreatedAt field in the returned CreateTransferOutput is correctly formatted as RFC3339 and matches the input future time.

Validation:
  This test checks that the presenter can handle future timestamps without modification, which might be important for scheduled or queued transfers in the system.
```

These test scenarios cover a range of normal operations, edge cases, and potential error conditions for the `Output` method of the `createTransferPresenter` struct. They aim to ensure that the method correctly transforms `domain.Transfer` objects into `usecase.CreateTransferOutput` objects under various conditions, maintaining data integrity and proper formatting throughout the process.
*/

// ********RoostGPT********
package presenter

import (
	"testing"
	"time"

	"github.com/gsabadini/go-clean-architecture/domain"
	"github.com/gsabadini/go-clean-architecture/usecase"
)

func TestcreateTransferPresenterOutput(t *testing.T) {
	tests := []struct {
		name     string
		transfer domain.Transfer
		want     usecase.CreateTransferOutput
	}{
		{
			name: "Valid Transfer Output",
			transfer: domain.Transfer{
				id:                   domain.TransferID("transfer123"),
				accountOriginID:      domain.AccountID("origin456"),
				accountDestinationID: domain.AccountID("dest789"),
				amount:               domain.Money(10050), // 100.50 * 100
				createdAt:            time.Date(2023, 5, 1, 10, 30, 0, 0, time.UTC),
			},
			want: usecase.CreateTransferOutput{
				ID:                   "transfer123",
				AccountOriginID:      "origin456",
				AccountDestinationID: "dest789",
				Amount:               100.50,
				CreatedAt:            "2023-05-01T10:30:00Z",
			},
		},
		{
			name: "Zero Amount Transfer",
			transfer: domain.Transfer{
				id:                   domain.TransferID("transfer000"),
				accountOriginID:      domain.AccountID("origin111"),
				accountDestinationID: domain.AccountID("dest222"),
				amount:               domain.Money(0),
				createdAt:            time.Date(2023, 5, 2, 15, 45, 0, 0, time.UTC),
			},
			want: usecase.CreateTransferOutput{
				ID:                   "transfer000",
				AccountOriginID:      "origin111",
				AccountDestinationID: "dest222",
				Amount:               0,
				CreatedAt:            "2023-05-02T15:45:00Z",
			},
		},
		{
			name: "Large Amount Transfer",
			transfer: domain.Transfer{
				id:                   domain.TransferID("transfer999"),
				accountOriginID:      domain.AccountID("origin888"),
				accountDestinationID: domain.AccountID("dest777"),
				amount:               domain.Money(99999999999), // 999999999.99 * 100
				createdAt:            time.Date(2023, 5, 3, 20, 0, 0, 0, time.UTC),
			},
			want: usecase.CreateTransferOutput{
				ID:                   "transfer999",
				AccountOriginID:      "origin888",
				AccountDestinationID: "dest777",
				Amount:               999999999.99,
				CreatedAt:            "2023-05-03T20:00:00Z",
			},
		},
		{
			name: "Transfer with Minimum Valid IDs",
			transfer: domain.Transfer{
				id:                   domain.TransferID("t1"),
				accountOriginID:      domain.AccountID("o1"),
				accountDestinationID: domain.AccountID("d1"),
				amount:               domain.Money(5025), // 50.25 * 100
				createdAt:            time.Date(2023, 5, 4, 9, 15, 0, 0, time.UTC),
			},
			want: usecase.CreateTransferOutput{
				ID:                   "t1",
				AccountOriginID:      "o1",
				AccountDestinationID: "d1",
				Amount:               50.25,
				CreatedAt:            "2023-05-04T09:15:00Z",
			},
		},
		{
			name: "Transfer with Maximum Valid IDs",
			transfer: domain.Transfer{
				id:                   domain.TransferID("transfer1234567890123456789012345678901234567890"),
				accountOriginID:      domain.AccountID("origin12345678901234567890123456789012345678901"),
				accountDestinationID: domain.AccountID("dest123456789012345678901234567890123456789012"),
				amount:               domain.Money(7575), // 75.75 * 100
				createdAt:            time.Date(2023, 5, 5, 14, 30, 0, 0, time.UTC),
			},
			want: usecase.CreateTransferOutput{
				ID:                   "transfer1234567890123456789012345678901234567890",
				AccountOriginID:      "origin12345678901234567890123456789012345678901",
				AccountDestinationID: "dest123456789012345678901234567890123456789012",
				Amount:               75.75,
				CreatedAt:            "2023-05-05T14:30:00Z",
			},
		},
		{
			name: "Transfer with Past Timestamp",
			transfer: domain.Transfer{
				id:                   domain.TransferID("transfer_past"),
				accountOriginID:      domain.AccountID("origin_past"),
				accountDestinationID: domain.AccountID("dest_past"),
				amount:               domain.Money(20000), // 200.00 * 100
				createdAt:            time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: usecase.CreateTransferOutput{
				ID:                   "transfer_past",
				AccountOriginID:      "origin_past",
				AccountDestinationID: "dest_past",
				Amount:               200.00,
				CreatedAt:            "2020-01-01T00:00:00Z",
			},
		},
		{
			name: "Transfer with Future Timestamp",
			transfer: domain.Transfer{
				id:                   domain.TransferID("transfer_future"),
				accountOriginID:      domain.AccountID("origin_future"),
				accountDestinationID: domain.AccountID("dest_future"),
				amount:               domain.Money(30000), // 300.00 * 100
				createdAt:            time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC),
			},
			want: usecase.CreateTransferOutput{
				ID:                   "transfer_future",
				AccountOriginID:      "origin_future",
				AccountDestinationID: "dest_future",
				Amount:               300.00,
				CreatedAt:            "2025-12-31T23:59:59Z",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := createTransferPresenter{}
			got := c.Output(tt.transfer)
			if got != tt.want {
				t.Errorf("createTransferPresenter.Output() = %v, want %v", got, tt.want)
			}
		})
	}
}
