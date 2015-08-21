# Yatzie - A telegram bot written in Golang

[![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/go-telegram-bot/yatzie?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

A customizable, pluggable Bot for telegram. Work in progress

Bot Commands
------------
<table>
  <thead>
    <tr>
      <td><strong>Name</strong></td>
      <td><strong>Description</strong></td>
      <td><strong>Usage</strong></td>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>9gag</td>
      <td>9GAG for telegram</td>
      <td>/gag: Send random image from 9gag</td>
    </tr>
    <tr>
      <td>boobs</td>
      <td>Gets a random boobs or butts pic</td>
      <td>/boobs: Get a boobs NSFW image. :underage:<br>/butts: Get a butts NSFW image. :underage:<br></td>
    </tr>
    <tr>
      <td>dogr</td>
      <td>Create a doge image with words.</td>
      <td>/doge BLA BLA BLA - Create a doge image with the words.</td>
    </tr>
    <tr>
      <td>google</td>
      <td>Searches Google and send results</td>
      <td>/search [terms]: Searches Google and send results<br>/imgsearch [terms]: Searches Google Images and send results</td>
    </tr>
    <tr>
      <td>hal</td>
      <td>Let the bot listen and learn</td>
      <td>say something to him or mention him (@bot) </td>
    </tr>
    <tr>
      <td>hello</td>
      <td>Let the bot says hello</td>
      <td>/hi to say hello</td>
    </tr>
    <tr>
      <td>help</td>
      <td>Help plugin. Get info from other plugins. </td>
      <td>/help: Show list of plugins</td>
    </tr>
    <tr>
      <td>imdb</td>
      <td>IMDB plugin for Telegram</td>
      <td>/imdb [movie]</td>
    </tr>
    <tr>
      <td>8ball</td>
      <td>Magic 8Ball</td>
      <td>/8ball [question]</td>
    </tr>
    <tr>
      <td>norris</td>
      <td>Send a random Chuck norris quote</td>
      <td>/norris</td>
    </tr>
    <tr>
      <td>xkcd</td>
      <td>Send comic images from xkcd</td>
      <td>/xkcd (id): Send an xkcd image and title. If not id, send the latest one<br></td>
    </tr>
  </tbody>
</table>

# Installation

```bash
go build github.com/go-telegram-bot/yatzie
```

# Configuration

The bot reads a json file for configuration, example:


```json
{
        "Token" : "XXXXXXX",
        "CommandPrefix" : "/"
}
```

# Run

```bash
$ yatzie -c yatzie.conf -l yatzie.log
```