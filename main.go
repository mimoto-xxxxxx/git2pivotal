package main

import (
	"git2pivotal/pivotal"
	"log"
)

func main() {
	token, err := getToken()
	if err != nil {
		log.Fatalln("Cannot get the token:", err)
	}
	if token == "" {
		log.Fatalln("The PivotalTracker API Token is empty")
	}

	project, err := getProject()
	if err != nil {
		log.Fatalln("Cannot get the project:", err)
	}
	if project == "" {
		log.Fatalln("The PivotalTracker Project ID is empty")
	}

	//この辺は省略可
	pre, err := getPreScriptTemplate()
	post, err := getPostScriptTemplate()

	//一番最新のコミットの本文だけを取得する
	commitMessage, err := getLatestCommit("%b")
	if err != nil {
		log.Fatalln("Cannot get the latest commit message:", err)
	}

	stories := pivotal.FindStoryFromString(commitMessage)
	if len(stories) == 0 {
		//コミット本文にストーリーIDらしきものがひとつもなかったので正常終了
		return
	}

	//テンプレートに従って値を取得
	var prescript, postscript string
	if prescript, err = getLatestCommit(pre); err != nil {
		log.Fatalln("Cannot get the postscript message:", err)
	}
	if postscript, err = getLatestCommit(post); err != nil {
		log.Fatalln("Cannot get the postscript message:", err)
	}

	//内容が空じゃないなら改行を加えておく
	if prescript != "" {
		prescript = prescript + "\n"
	}
	if postscript != "" {
		postscript = "\n" + postscript
	}

	//ノート本文を組み立て
	//prescript + ストーリーIDを取り除いた本文 + postscript
	commitNoStory := prescript + pivotal.RemoveStoryFromString(commitMessage) + postscript

	//同じ本文を全てのストーリーのノートとして送信
	for _, story := range stories {
		if err = pivotal.AddNote(token, project, story.Id, commitNoStory); err != nil {
			//途中でエラーが起きたら報告はするけど処理は続行
			log.Printf("add note error(Project:%s / Story:%s): %v", project, story.Id, err)
		}
	}
}
