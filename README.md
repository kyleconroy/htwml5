Power your Twilio applications using HTML5. 

## Text to speech

```html
<html>
  <body>
    <p>Hello world</p>
  </body>
</html>
```

You can change the voice and language of the spoken text.

```html
<html>
  <body>
    <p lang="fr" data-voice="woman">Hello world</p>
  </body>
</html>
```

If you'd like to insert a pause, use an `<hr>` element.

```html
<html>
  <body>
    <p>Hello World</p>
    <hr data-length=10 />
    <p>How are you?</p>
  </body>
</html>
```

## Phone Menus

Using the little known [menu element](), you can easily build phone trees, no
server-side logic required. Links must have the `digit` data attribute for menu
routing to work. This value can be an integer.

```html
<html>
  <body>
    <menu type="list">
      <li><a href="/monkey.html" data-digit=1>For a monkey, press 1</a></li>
      <li><a href="/gorilla.html" data-digit=2>For a gorilla, press 2</a></li>
    </menu>
  </body>
</html>
```

## URI schemes

URI schemes allow users to directly connect to phones, conferences, and clients

```html
<a href="tel:+15555551234" data-digit=1>To dial a phone, press 1</a>
```

```html
<a href="conference:foo" data-digit=1>To dial a conference, press 1</a>
```

```html
<a href="client:foo" data-digit=1>To dial a browser, press 1</a>
```

## Audio

```html
<html>
  <body>
    <audio src="http://demo.twilio.com/hellomonkey/monkey.mp3" controls></audio>
  </body>
</html>
```

```html
<html>
  <body>
    <audio> 
      <source src="http://demo.twilio.com/hellomonkey/monkey.mp3">
      <p>This text will be read if the above url fails to load</p>
    </audio>
  </body>
</html>
```


## Forms

Forms support text, hidden, and select inputs. You can only have one text or
select input per form. Label text is also read

In the future, we may support multiple inputs via some magic.

```html
<html>
  <body>
    <form action="/keypads/process" method="POST">
      <label for="foo">Please enter your zipcode</label>
      <input name="foo" type="text">
    </form> 
  </body>
</html>
```

### Validation

You currently must to validation server-side. In the future, we may have
default error messages if validation attributes are provided

## Recordings

Recording speech works in a similar fashion, just specify the speech attribute

```html
<html>
  <body>
    <form action="/keypads/process" method="POST">
      <input type="text" speech>
    </form> 
  </body>
</html>
```

## Well-formed HTML

The above examples are concise, but not entirely correct. To ensure validation
and proper support in the future, make sure you declare a DOCTYPE,
character-encoding, and langauge.

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
  </head>
  <body>
    ... content here ...
  </body>
</html>
```
