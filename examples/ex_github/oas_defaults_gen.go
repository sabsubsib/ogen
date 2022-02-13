// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/bits"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = bits.LeadingZeros64
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
	_ = codes.Unset
)

// setDefaults set default value of fields.
func (s *ActionsCreateSelfHostedRunnerGroupForOrgReq) setDefaults() {
	{
		val := ActionsCreateSelfHostedRunnerGroupForOrgReqVisibility("all")

		s.Visibility.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *ChecksSetSuitesPreferencesReqAutoTriggerChecksItem) setDefaults() {
	{
		val := bool(true)

		s.Setting = val
	}
}

// setDefaults set default value of fields.
func (s *EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseReq) setDefaults() {
	{
		val := EnterpriseAdminUpdateSelfHostedRunnerGroupForEnterpriseReqVisibility("all")

		s.Visibility.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *FullRepository) setDefaults() {
	{
		val := bool(true)

		s.AnonymousAccessEnabled.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *GitCreateBlobReq) setDefaults() {
	{
		val := string("utf-8")

		s.Encoding.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *IssuesCreateMilestoneReq) setDefaults() {
	{
		val := IssuesCreateMilestoneReqState("open")

		s.State.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *IssuesUpdateMilestoneReq) setDefaults() {
	{
		val := IssuesUpdateMilestoneReqState("open")

		s.State.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *Milestone) setDefaults() {
	{
		val := MilestoneState("open")

		s.State = val
	}
}

// setDefaults set default value of fields.
func (s *NullableMilestone) setDefaults() {
	{
		val := NullableMilestoneState("open")

		s.State = val
	}
}

// setDefaults set default value of fields.
func (s *NullableRepository) setDefaults() {
	{
		val := bool(true)

		s.HasIssues = val
	}
	{
		val := bool(true)

		s.HasProjects = val
	}
	{
		val := bool(true)

		s.HasWiki = val
	}
	{
		val := bool(true)

		s.HasDownloads = val
	}
	{
		val := string("public")

		s.Visibility.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowRebaseMerge.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowSquashMerge.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowMergeCommit.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *OrgsCreateInvitationReq) setDefaults() {
	{
		val := OrgsCreateInvitationReqRole("direct_member")

		s.Role.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *OrgsCreateWebhookReq) setDefaults() {
	{
		s.Events = make([]string, 1)
		{
			val := string("push")

			s.Events[0] = val
		}
	}
	{
		val := bool(true)

		s.Active.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *OrgsSetMembershipForUserReq) setDefaults() {
	{
		val := OrgsSetMembershipForUserReqRole("member")

		s.Role.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *OrgsUpdateWebhookReq) setDefaults() {
	{
		s.Events = make([]string, 1)
		{
			val := string("push")

			s.Events[0] = val
		}
	}
	{
		val := bool(true)

		s.Active.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *ProjectsAddCollaboratorReq) setDefaults() {
	{
		val := ProjectsAddCollaboratorReqPermission("write")

		s.Permission.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *PullRequestReviewComment) setDefaults() {
	{
		val := PullRequestReviewCommentStartSide("RIGHT")

		s.StartSide.SetTo(val)
	}
	{
		val := PullRequestReviewCommentSide("RIGHT")

		s.Side.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *ReposAddCollaboratorReq) setDefaults() {
	{
		val := ReposAddCollaboratorReqPermission("push")

		s.Permission.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *ReposCreateCommitStatusReq) setDefaults() {
	{
		val := string("default")

		s.Context.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *ReposCreateDeploymentReq) setDefaults() {
	{
		val := string("deploy")

		s.Task.SetTo(val)
	}
	{
		val := bool(true)

		s.AutoMerge.SetTo(val)
	}
	{
		val := string("production")

		s.Environment.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *ReposCreateForAuthenticatedUserReq) setDefaults() {
	{
		val := bool(true)

		s.HasIssues.SetTo(val)
	}
	{
		val := bool(true)

		s.HasProjects.SetTo(val)
	}
	{
		val := bool(true)

		s.HasWiki.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowSquashMerge.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowMergeCommit.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowRebaseMerge.SetTo(val)
	}
	{
		val := bool(true)

		s.HasDownloads.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *ReposCreateInOrgReq) setDefaults() {
	{
		val := bool(true)

		s.HasIssues.SetTo(val)
	}
	{
		val := bool(true)

		s.HasProjects.SetTo(val)
	}
	{
		val := bool(true)

		s.HasWiki.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowSquashMerge.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowMergeCommit.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowRebaseMerge.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *ReposCreatePagesSiteReqSource) setDefaults() {
	{
		val := ReposCreatePagesSiteReqSourcePath("/")

		s.Path.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *ReposCreateWebhookReq) setDefaults() {
	{
		s.Events = make([]string, 1)
		{
			val := string("push")

			s.Events[0] = val
		}
	}
	{
		val := bool(true)

		s.Active.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *ReposUpdateReq) setDefaults() {
	{
		val := bool(true)

		s.HasIssues.SetTo(val)
	}
	{
		val := bool(true)

		s.HasProjects.SetTo(val)
	}
	{
		val := bool(true)

		s.HasWiki.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowSquashMerge.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowMergeCommit.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowRebaseMerge.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *ReposUpdateWebhookReq) setDefaults() {
	{
		s.Events = make([]string, 1)
		{
			val := string("push")

			s.Events[0] = val
		}
	}
	{
		val := bool(true)

		s.Active.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *Repository) setDefaults() {
	{
		val := bool(true)

		s.HasIssues = val
	}
	{
		val := bool(true)

		s.HasProjects = val
	}
	{
		val := bool(true)

		s.HasWiki = val
	}
	{
		val := bool(true)

		s.HasDownloads = val
	}
	{
		val := string("public")

		s.Visibility.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowRebaseMerge.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowSquashMerge.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowMergeCommit.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *ReviewComment) setDefaults() {
	{
		val := ReviewCommentSide("RIGHT")

		s.Side.SetTo(val)
	}
	{
		val := ReviewCommentStartSide("RIGHT")

		s.StartSide.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *TeamMembership) setDefaults() {
	{
		val := TeamMembershipRole("member")

		s.Role = val
	}
}

// setDefaults set default value of fields.
func (s *TeamRepository) setDefaults() {
	{
		val := bool(true)

		s.HasIssues = val
	}
	{
		val := bool(true)

		s.HasProjects = val
	}
	{
		val := bool(true)

		s.HasWiki = val
	}
	{
		val := bool(true)

		s.HasDownloads = val
	}
	{
		val := string("public")

		s.Visibility.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowRebaseMerge.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowSquashMerge.SetTo(val)
	}
	{
		val := bool(true)

		s.AllowMergeCommit.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *TeamsAddOrUpdateMembershipForUserInOrgReq) setDefaults() {
	{
		val := TeamsAddOrUpdateMembershipForUserInOrgReqRole("member")

		s.Role.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *TeamsAddOrUpdateMembershipForUserLegacyReq) setDefaults() {
	{
		val := TeamsAddOrUpdateMembershipForUserLegacyReqRole("member")

		s.Role.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *TeamsCreateReq) setDefaults() {
	{
		val := TeamsCreateReqPermission("pull")

		s.Permission.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *TeamsUpdateInOrgReq) setDefaults() {
	{
		val := TeamsUpdateInOrgReqPermission("pull")

		s.Permission.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *TeamsUpdateLegacyReq) setDefaults() {
	{
		val := TeamsUpdateLegacyReqPermission("pull")

		s.Permission.SetTo(val)
	}
}