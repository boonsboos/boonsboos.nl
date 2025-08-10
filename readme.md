# boonsboos.nl

custom blog system for my personal site

### features
- write your posts in markdown
- easy to style with [tailwind](https://tailwindcss.com)
- style your code tags yourself with [chroma](https://github.com/alecthomas/chroma)

### requirements
- go (>=1.21)
- node (>= 20)

### building
```shell
npm install
npx @tailwindcss/cli -i ./base.tailwind.css -o ./static/s.css
export GIN_MODE=release; go run .
```

### is this for you?
short answer: not really

long answer: I made this site generator for my own personal website and therefore it only has the things I need. If you want to add a feature, feel free to fork the project and put your own twist on it.