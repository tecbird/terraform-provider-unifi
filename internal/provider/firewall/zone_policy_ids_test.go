package firewall

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestIntsToStrings(t *testing.T) {
	got := intsToStrings([]int{589888, 1310917})
	want := []string{"589888", "1310917"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestStringListToInts(t *testing.T) {
	ctx := context.Background()
	list := types.ListValueMust(types.StringType, []attr.Value{types.StringValue("589888"), types.StringValue("1310917")})

	var got []int
	diags := stringListToInts(ctx, list, &got, "app_ids")
	if diags.HasError() {
		t.Fatalf("unexpected error: %v", diags)
	}
	if !reflect.DeepEqual(got, []int{589888, 1310917}) {
		t.Fatalf("got %v", got)
	}

	bad := types.ListValueMust(types.StringType, []attr.Value{types.StringValue("abc")})
	if !stringListToInts(ctx, bad, &got, "app_ids").HasError() {
		t.Fatal("expected an error for a non-numeric ID")
	}
}
