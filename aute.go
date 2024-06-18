import (
	"context"
	"fmt"
	"io"

	iam "google.golang.org/api/iam/v1"
)

// approveDelegation approves a service account delegation token request.
func approveDelegation(w io.Writer, projectID, delegationID string) (*iam.ServiceAccountDelegation, error) {
	ctx := context.Background()
	service, err := iam.NewService(ctx)
	if err != nil {
		return nil, fmt.Errorf("iam.NewService: %v", err)
	}

	delegationsService := service.Projects.ServiceAccounts.Delegations

	req := &iam.ApproveDelegationRequest{}
	resp, err := delegationsService.Approve(projectID, delegationID, req).Do()
	if err != nil {
		return nil, fmt.Errorf("ApproveDelegation: %v", err)
	}

	fmt.Fprintf(w, "Delegation approved: %v", resp.Name)
	return resp, nil
}
  
