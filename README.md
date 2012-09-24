Power your Twilio applications using HTML5. 

## Text to speech

```html
<html>
  <body>
    <p>Hello world</p>
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



