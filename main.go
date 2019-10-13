package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"context"
	"log"
	"os"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

func ParseJson() (string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var lineByte = []byte(line)
	var v interface{}
	var w interface{}
	err := json.Unmarshal(lineByte, &v)
	content := v.([]interface{})[0].(map[string]interface{})["content"].(string)
	contentByte := []byte(content)
	err = json.Unmarshal(contentByte, w)
	comment := w.([]interface{})[0].(map[string]interface{})["comment"].(string)
	commentByte := []byte(comment)
	if err != nil {
		log.Printf("%s", err)
	}
	if w.([]interface{})[0] == "comment" {
		file, err := os.Create("log.txt")
		if err != nil {
			log.Printf("%s", err)
		}
		defer file.Close()
		_, _ = file.Write(commentByte)
	}
	return comment
}

func speach(text string) {
       // Instantiates a client.
       ctx := context.Background()

       client, err := texttospeech.NewClient(ctx)
       if err != nil {
               log.Fatal(err)
       }

       // Perform the text-to-speech request on the text input with the selected
       // voice parameters and audio file type.
       req := texttospeechpb.SynthesizeSpeechRequest{
               // Set the text input to be synthesized.
               Input: &texttospeechpb.SynthesisInput{
                       InputSource: &texttospeechpb.SynthesisInput_Text{Text: "こんにちは"},
               },
               // Build the voice request, select the language code ("en-US") and the SSML
               // voice gender ("neutral").
               Voice: &texttospeechpb.VoiceSelectionParams{
                       LanguageCode: "ja-JP",
                       SsmlGender:   texttospeechpb.SsmlVoiceGender_NEUTRAL,
               },
               // Select the type of audio file you want returned.
               AudioConfig: &texttospeechpb.AudioConfig{
                       AudioEncoding: texttospeechpb.AudioEncoding_MP3,
               },
       }
       resp, err := client.SynthesizeSpeech(ctx, &req)
       if err != nil {
               log.Fatal(err)
       }
       // The resp's AudioContent is binary.
       filename := "output.mp3"
       err = ioutil.WriteFile(filename, resp.AudioContent, 0644)
       if err != nil {
               log.Fatal(err)
       }
       fmt.Printf("Audio content written to file: %v\n", filename)
}

func main()  {
	for {
		text := ParseJson()
		speach(text)
	}
}
