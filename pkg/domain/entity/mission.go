package entity

type MissionCondition string

const (
	MissionConditionNone                 MissionCondition = "None"                 // なし
	MissionConditionLogin                MissionCondition = "Login"                // [value1]日ログインする
	MissionConditionQuestPlay            MissionCondition = "QuestPlay"            // いずれかのクエストを[value1]回プレイする
	MissionConditionQuestRankS           MissionCondition = "QuestRankS"           // いずれかのクエストでSランクを[value1]回獲得する
	MissionConditionMainQuestClear       MissionCondition = "MainQuestClear"       // メインクエストを[value1]回クリアする※[value2]でID指定
	MissionConditionSubQuestClear        MissionCondition = "SubQuestClear"        // いずれかのサブクエストを[value1]回クリアする
	MissionConditionMatchPiece           MissionCondition = "MatchPiece"           // ピースを[value1]個消す
	MissionConditionMatchCombo           MissionCondition = "MatchCombo"           // 合計[value1]コンボ達成する
	MissionConditionMatchMaxCombo        MissionCondition = "MatchMaxCombo"        // 最大[value1]コンボを達成する
	MissionConditionEnemyDefeat          MissionCondition = "EnemyDefeat"          // 合計[value1]体の敵を倒す※[value2]でID指定
	MissionConditionFreePresent          MissionCondition = "FreePresent"          // 配信者に[value1]個ハピネスボックスを贈る
	MissionConditionCharacterEnhance     MissionCondition = "CharacterEnhance"     // いずれかのキャラクターを[value1]までレベルアップしよう！※[value2]でID指定
	MissionConditionCharacterAwake       MissionCondition = "CharacterAwake"       // いずれかのキャラクターを[value1]まで覚醒しよう！※[value2]でID指定
	MissionConditionCompleteMissionGroup MissionCondition = "CompleteMissionGroup" // 同じグループ内のミッションをすべてクリアしよう
	MissionConditionTotalLogin           MissionCondition = "TotalLogin"           // 累計[value1]日間ログインする
	MissionConditionTotalQuestRankS      MissionCondition = "TotalQuestRankS"      // 累計クエストでSランクを[value1]個取る
	MissionConditionTotalQuestClear      MissionCondition = "TotalQuestClear"      // 累計[value1]回いずれかのクエストクリア※プレイではない
	MissionConditionTotalFreePresent     MissionCondition = "TotalFreePresent"     // 累計[value1]回ハピネスボックスを贈る
)

type Mission struct {
	SeedBase        `yaml:",inline"`
	MissionGroupID  uint             `yaml:"groupId"`
	MissionGroup    MissionGroup     `gorm:"foreignKey:MissionGroupID"`
	Text            string           `yaml:"text"`
	IsLive          bool             `yaml:"isLive"`
	Condition       MissionCondition `yaml:"condition"`
	ConditionValue1 int              `yaml:"conditionValue1"`
	ConditionValue2 int              `yaml:"conditionValue2"`
	RewardGroupID   uint             `yaml:"rewardGroupId"`
}
