package processtree

import (
	"reflect"
	"testing"

	"example.com/security/internal/models"
)

func TestPrune(t *testing.T) {
	tests := []struct {
		name             string
		processesCreated []string
		nodes            []models.ProcessTree
		want             []models.ProcessTree
	}{
		{
			name:             "空のNodeは空を返す",
			processesCreated: []string{"cmd.exe"},
			nodes:            []models.ProcessTree{},
			want:             []models.ProcessTree{},
		},
		{
			name:             "processes_createdに存在しない単一Nodeは除外される",
			processesCreated: []string{"cmd.exe"},
			nodes: []models.ProcessTree{
				{Name: "notepad.exe", ProcessID: "1"},
			},
			want: []models.ProcessTree{},
		},
		{
			name:             "processes_createdに存在する単一Nodeは保持される",
			processesCreated: []string{"cmd.exe"},
			nodes: []models.ProcessTree{
				{Name: "cmd.exe", ProcessID: "1"},
			},
			want: []models.ProcessTree{
				{Name: "cmd.exe", ProcessID: "1", Children: []models.ProcessTree{}},
			},
		},
		{
			name:             "親がマッチしなくても子がマッチすれば親は保持され、子のみ残る",
			processesCreated: []string{"powershell.exe"},
			nodes: []models.ProcessTree{
				{
					Name:      "explorer.exe",
					ProcessID: "1",
					Children: []models.ProcessTree{
						{Name: "powershell.exe", ProcessID: "2"},
					},
				},
			},
			want: []models.ProcessTree{
				{
					Name:      "explorer.exe",
					ProcessID: "1",
					Children: []models.ProcessTree{
						{Name: "powershell.exe", ProcessID: "2", Children: []models.ProcessTree{}},
					},
				},
			},
		},
		{
			name:             "親がマッチし子がマッチしない場合、親は残るが該当しない子は除外される",
			processesCreated: []string{"cmd.exe"},
			nodes: []models.ProcessTree{
				{
					Name:      "cmd.exe",
					ProcessID: "1",
					Children: []models.ProcessTree{
						{Name: "notepad.exe", ProcessID: "2"},
					},
				},
			},
			want: []models.ProcessTree{
				{
					Name:      "cmd.exe",
					ProcessID: "1",
					Children:  []models.ProcessTree{},
				},
			},
		},
		{
			name:             "兄弟Nodeのうちマッチするものだけが残る",
			processesCreated: []string{"cmd.exe", "powershell.exe"},
			nodes: []models.ProcessTree{
				{Name: "cmd.exe", ProcessID: "1"},
				{Name: "notepad.exe", ProcessID: "2"},
				{Name: "powershell.exe", ProcessID: "3"},
			},
			want: []models.ProcessTree{
				{Name: "cmd.exe", ProcessID: "1", Children: []models.ProcessTree{}},
				{Name: "powershell.exe", ProcessID: "3", Children: []models.ProcessTree{}},
			},
		},
		{
			name:             "多段階(孫)でマッチする場合は経路上の全Nodeが保持される",
			processesCreated: []string{"rundll32.exe"},
			nodes: []models.ProcessTree{
				{
					Name:      "explorer.exe",
					ProcessID: "1",
					Children: []models.ProcessTree{
						{
							Name:      "cmd.exe",
							ProcessID: "2",
							Children: []models.ProcessTree{
								{Name: "rundll32.exe", ProcessID: "3"},
							},
						},
					},
				},
			},
			want: []models.ProcessTree{
				{
					Name:      "explorer.exe",
					ProcessID: "1",
					Children: []models.ProcessTree{
						{
							Name:      "cmd.exe",
							ProcessID: "2",
							Children: []models.ProcessTree{
								{Name: "rundll32.exe", ProcessID: "3", Children: []models.ProcessTree{}},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Prune(tt.processesCreated, tt.nodes)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prune() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestBuild(t *testing.T) {
	tests := []struct {
		name  string
		nodes []models.ProcessTree
		want  []models.ProcessTree
	}{
		{
			name:  "空のNodeは空を返す",
			nodes: []models.ProcessTree{},
			want:  []models.ProcessTree{},
		},
		{
			name: "単一Nodeはそのまま変換される",
			nodes: []models.ProcessTree{
				{Name: "cmd.exe", ProcessID: "1"},
			},
			want: []models.ProcessTree{
				{Name: "cmd.exe", ProcessID: "1", Children: []models.ProcessTree{}},
			},
		},
		{
			name: "同名かつ子構造が同一(空)のトップレベルNodeは1つに重複排除される",
			nodes: []models.ProcessTree{
				{Name: "cmd.exe", ProcessID: "1"},
				{Name: "cmd.exe", ProcessID: "2"},
			},
			want: []models.ProcessTree{
				{Name: "cmd.exe", ProcessID: "1", Children: []models.ProcessTree{}},
			},
		},
		{
			name: "同名でも子構造が異なる場合は別Nodeとして保持される",
			nodes: []models.ProcessTree{
				{
					Name:      "cmd.exe",
					ProcessID: "1",
					Children: []models.ProcessTree{
						{Name: "powershell.exe", ProcessID: "10"},
					},
				},
				{
					Name:      "cmd.exe",
					ProcessID: "2",
					Children: []models.ProcessTree{
						{Name: "notepad.exe", ProcessID: "20"},
					},
				},
			},
			want: []models.ProcessTree{
				{
					Name:      "cmd.exe",
					ProcessID: "1",
					Children: []models.ProcessTree{
						{Name: "powershell.exe", ProcessID: "10", Children: []models.ProcessTree{}},
					},
				},
				{
					Name:      "cmd.exe",
					ProcessID: "2",
					Children: []models.ProcessTree{
						{Name: "notepad.exe", ProcessID: "20", Children: []models.ProcessTree{}},
					},
				},
			},
		},
		{
			name: "深い階層(子)の重複も再帰的に排除される",
			nodes: []models.ProcessTree{
				{
					Name:      "explorer.exe",
					ProcessID: "1",
					Children: []models.ProcessTree{
						{Name: "cmd.exe", ProcessID: "10"},
						{Name: "cmd.exe", ProcessID: "11"},
					},
				},
			},
			want: []models.ProcessTree{
				{
					Name:      "explorer.exe",
					ProcessID: "1",
					Children: []models.ProcessTree{
						{Name: "cmd.exe", ProcessID: "10", Children: []models.ProcessTree{}},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Build(tt.nodes)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
