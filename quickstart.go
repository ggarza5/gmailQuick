package main

import (
        "encoding/json"
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
        "os"
        "golang.org/x/net/context"
        "golang.org/x/oauth2"
        "golang.org/x/oauth2/google"
        "google.golang.org/api/gmail/v1"
        "strings"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
        // The file token.json stores the user's access and refresh tokens, and is
        // created automatically when the authorization flow completes for the first
        // time.
        tokFile := "token.json"
        tok, err := tokenFromFile(tokFile)
        if err != nil {
                tok = getTokenFromWeb(config)
                saveToken(tokFile, tok)
        }
        return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
        authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
        fmt.Printf("Go to the following link in your browser then type the "+
                "authorization code: \n%v\n", authURL)

        var authCode string
        if _, err := fmt.Scan(&authCode); err != nil {
                log.Fatalf("Unable to read authorization code: %v", err)
        }

        tok, err := config.Exchange(context.TODO(), authCode)
        if err != nil {
                log.Fatalf("Unable to retrieve token from web: %v", err)
        }
        return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
        f, err := os.Open(file)
        if err != nil {
                return nil, err
        }
        defer f.Close()
        tok := &oauth2.Token{}
        err = json.NewDecoder(f).Decode(tok)
        return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
        fmt.Printf("Saving credential file to: %s\n", path)
        f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
        if err != nil {
                log.Fatalf("Unable to cache oauth token: %v", err)
        }
        defer f.Close()
        json.NewEncoder(f).Encode(token)
}

func getPcEmail(srv *gmail.Client) *gmail.Message {
        user := "me"
        
}

func getPageOfEmails(srv *gmail.Client) {
        user := "me"
        msgs, e2  := srv.Users.Messages.List(user).Do()

}

func isPcEmail(email string) bool {
        words := strings.Split(email, ' ') 
        for _, w := range words {
                if word == "PatternCast" {
                        return true
                }
        }
        return false
}

func main() {
        b, err := ioutil.ReadFile("credentials.json")
        if err != nil {
                log.Fatalf("Unable to read client secret file: %v", err)
        }

        // If modifying these scopes, delete your previously saved token.json.
        config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
        if err != nil {
                log.Fatalf("Unable to parse client secret file to config: %v", err)
        }
        client := getClient(config)

        srv, err := gmail.New(client)
        if err != nil {
                log.Fatalf("Unable to retrieve Gmail client: %v", err)
        }

        user := "me"
        r, err := srv.Users.Labels.List(user).Do()
        if err != nil {
                log.Fatalf("Unable to retrieve labels: %v", err)
        }
        if len(r.Labels) == 0 {
                fmt.Println("No labels found.")
                return
        }
        println("before labels")
        fmt.Println("Labels:")
        // for _, l := range r.Labels {
        //         fmt.Printf("- %s\n", l.Name)
        // }

        msgs, e2  := srv.Users.Messages.List(user).Do()
        if e2 != nil {
                log.Fatalf("Unable to retrieve Messages: %v", e2)
        }
        //println(m.Id)
        for _, m := range msgs.Messages {
                // ee, mm := srv.UsersMessages println(m.Id)
                mm, _ := srv.Users.Messages.Get(user, m.Id).Do()
                // if ee != nil {
                //         println("Ee not nil")
                //         println(*ee)
                // }
                println(mm.Snippet)
        }
}