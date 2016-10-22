package main

import (
	"html/template"
	"io"
	"net/http"
	//"os"
)

// 1. add news struct;	 2. add arrays;  3. add http; 4. clean others   5. data type * 3;   6. add styles
// 7. add html

type Gopher string

type News struct {
	Title   string
	Desc    string
	Img     string
	Url     string
	Author  string
	Content string
}

type HomeData struct {
	User     string
	Carousel [3]News
	ListNews [3]News
	Channel1 [3]News
	Channel2 [3]News
}

func hello(w http.ResponseWriter, r *http.Request) {

	s := "<html><body><style>.news-block{width:500px;background-color: powderblue;}</style></body></html>"
	io.WriteString(w, s)

	t := template.New("home")
	t, _ = t.Parse(hometmpl)
	t, _ = t.Parse(carouseltmpl)
	t, _ = t.Parse(listtmpl)
	t, _ = t.Parse(newstmpl)
	t, _ = t.Parse(newstmpl2)

	var newsdata, newsdataL, newsdata1, newsdata2 [3]News

	newsdata[2] = News{
		"Footballer dies after mass drug overdose on Gold Coast",
		"Police issue a warning over the dangers of synthetic drugs after Victorian footballer Riki Stephens,… ",
		"http://www.abc.net.au/news/image/7955280-3x2-460x307.jpg",
		"http://www.abc.net.au/news/2016-10-21/riki-stephens-gold-coast-drug-overdose-victorian-footballer-dies/7953486",
		"ABC News",
		"......1st the news content N/A. yet ......",
	}
	newsdata[1] = News{
		"Perth father charged with murder of his two children",
		"A 35-year-old Yanchep man under police guard in hospital is charged with two counts of murder over…",
		"http://www.abc.net.au/news/image/7955030-3x2-460x307.jpg",
		"http://www.abc.net.au/news/2016-10-22/perth-man-charged-with-murder-of-two-children/7957008",
		"ABC News",
		"......2nd the news content N/A. yet ......",
	}
	newsdata[0] = News{
		"Turnbull defeats Abbott in battle for party preselection&nbsp;reform",
		"Former prime minister Tony Abbott's push to change the NSW Liberal Party's preselection processes is… ",
		"http://www.abc.net.au/news/image/7952692-3x2-460x307.jpg",
		"http://www.abc.net.au/news/2016-10-22/nsw-liberals-back-turnbulls-preselection-reform-proposal/7957042",
		"ABC News",
		"Former prime minister Tony Abbott's push to change the NSW Liberal Party's preselection processes is defeated with the state council instead unanimously backing Malcolm Turnbull's plan.",
	}

	newsdataL[0] = News{
		"First New US Nuclear Reactor In 20 Years Goes Live",
		"The Tennessee Valley Authority is celebrating an event 43 years in the making: the completion of the Watts Bar Nuclear Plant. ",
		"",
		"https://hardware.slashdot.org/story/16/10/20/2130207/first-new-us-nuclear-reactor-in-20-years-goes-live?sbsrc=md",
		"",
		"",
	}
	newsdataL[1] = News{
		`Elon Musk: Negative Media Coverage of Autonomous Vehicles Could be 'Killing people' `,
		`On the sidelines of the Tesla announcements, CEO Elon Musk accused media of "killing people" by dissuading consumers from using an autonomous vehicle.`,
		``,
		`https://news.slashdot.org/story/16/10/20/196211/elon-musk-negative-media-coverage-of-autonomous-vehicles-could-be-killing-people?sbsrc=md`,
		``,
		``,
	}

	newsdataL[2] = News{
		"AI Platform Assesses Trump's and Clinton's Emotional Intelligence",
		"FastCompany got an exclusive look at how Hillary Clinton and Donald Trump stacked up in terms of their emotional intelligence when analyzed by HireVue's artificial intelligence platform. ",
		"",
		"https://politics.slashdot.org/story/16/10/21/219237/ai-platform-assesses-trumps-and-clintons-emotional-intelligence",
		"",
		"",
	}

	newsdata1[0] = News{
		"So long Salesforce, hello Softbank",
		"Will Twitter be acquired by a larger corporation for upward of $20 billion later this month? We’ve compiled the essential roundup of all the rumors surrounding the tech acquisition that has everyone talking.",
		"http://icdn8.digitaltrends.com/image/twitter-phone-photo-372x186-c.jpg",
		"http://www.digitaltrends.com/social-media/twitter-takeoever-roundup/",
		"Saqib Shah",
		"Will Twitter be acquired by a larger corporation for upward of $20 billion later this month? We’ve compiled the essential roundup of all the rumors surrounding the tech acquisition that has everyone talking.",
	}

	newsdata1[1] = News{
		"Show and Tell: Facebook Live ad campaign features tutorials and crowd-sourced footage",
		"Facebook is launching a large-scale ad campaign on the streets of the U.S. and UK in a bid to get more people to start streaming using its Live Video service. The ads feature crowd-sourced video clips and live-streaming tutorials.",
		"http://icdn2.digitaltrends.com/image/facebook-live-tutorial-1-372x186-c.jpg",
		"http://www.digitaltrends.com/social-media/facebook-live-ads/",
		"Saqib Shah",
		"Facebook is launching a large-scale ad campaign on the streets of the U.S. and UK in a bid to get more people to start streaming using its Live Video service. The ads feature crowd-sourced video clips and live-streaming tutorials.",
	}

	newsdata1[2] = News{
		`Facebook to permit NSFW content deemed to be of "public interest"`,
		`Facebook will start allowing more “newsworthy” content deemed to be of "public interest" on to its site even if those items break its Community Standards, pertaining to nudity and graphic violence.`,
		`http://icdn3.digitaltrends.com/image/facebook-headquarters-zuckerberg-372x186-c.jpg?ver=1`,
		`http://www.digitaltrends.com/social-media/facebook-nsfw-news-policy/`,
		`Saqib Shah`,
		`Facebook will start allowing more “newsworthy” content deemed to be of "public interest" on to its site even if those items break its Community Standards, pertaining to nudity and graphic violence.`,
	}

	newsdata2[0] = News{
		`Here’s your first look at Lucid Motors’ 900-horsepower, 300-mile Tesla fighter`,
		`California-based Lucid Motors has released several teaser images of its upcoming electric production car. Lucid Motors is led by former Tesla VP and Model S Chief Engineer Peter Rawlinson.`,
		`http://icdn2.digitaltrends.com/image/intro_lucid_blog_desert-372x186-c.png`,
		`http://www.digitaltrends.com/cars/lucid-motors-previews-production-car-pictures-specs-tesla/`,
		`Andrew Hard`,
		`California-based Lucid Motors has released several teaser images of its upcoming electric production car. Lucid Motors is led by former Tesla VP and Model S Chief Engineer Peter Rawlinson.`,
	}
	newsdata2[1] = News{
		`The FAA says to keep drones out of disaster relief areas, major sporting events`,
		`The Federal Aviation Administration has some stiff penalties in place for piloting drones into the wrong areas, such as where disaster relief efforts are underway and major sporting events.`,
		`http://icdn2.digitaltrends.com/image/quadcopter-3-372x186-c.jpg`,
		`http://www.digitaltrends.com/cool-tech/faa-drone-regulations-relief-efforts-sporting-events/`,
		`Mark Coppock`,
		`The Federal Aviation Administration has some stiff penalties in place for piloting drones into the wrong areas, such as where disaster relief efforts are underway and major sporting events.`,
	}
	newsdata2[2] = News{
		`Update: Morning attack on DNS provider resumes, internet burns`,
		`A cyberattack on Dyn, a major internet management company, has left a significant portion of the web in shambles, with users reporting issues with popular sites like Twitter, Spotify, SoundCloud, Airbnb, and more.`,
		`http://icdn2.digitaltrends.com/image/data_center_feat-372x186-c.jpg`,
		`http://www.digitaltrends.com/computing/dyn-ddos-attack/`,
		`Lulu Chang`,
		`A cyberattack on Dyn, a major internet management company, has left a significant portion of the web in shambles, with users reporting issues with popular sites like Twitter, Spotify, SoundCloud, Airbnb, and more.`,
	}

	/*
		newsdata[2] = News{
			"Wonder Woman named honorary UN ambassador; not everyone is happy about it",
			"From her origin to her sexual orientation right down to her latest toned-down outfit, arguably no figure … ",
			"http://www.theage.com.au/content/dam/images/g/q/k/r/v/a/image.related.articleLeadwide.620x349.gs891p.png/1477097486324.jpg",
			"http://www.theage.com.au/world/wonder-woman-named-honorary-un-ambassador-not-everyone-is-happy-about-it-20161022-gs891p.html",
			"Michael Cavna",
			"From her origin to her sexual orientation right down to her latest toned-down outfit, arguably no figure in all of superhero comics is more saddled with political scrutiny great and small than Wonder Woman. She projects strength with compassion; onto her, the world has projected so much else.",
		}
	*/
	data := &HomeData{
		User:     "Stephen Li",
		Carousel: newsdata,
		ListNews: newsdataL,
		Channel1: newsdata1,
		Channel2: newsdata2,
	}

	t.ExecuteTemplate(w, "home", data)

}

func main() {

	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}

/*
const hometmpl = `
{{define "home"}}
	{{template "carousel".Carousel}}
	{{template "listnews".ListNews}}
	{{template "news".Channel1}}
	{{template "news".Channel2}}
{{end}}
`
*/

const carouseltmpl = `
{{define "carousel"}}
      <div class="carousel-inner" role="listbox">
{{range $i,$e :=.}} {{if .Title}}      
    {{if not $i}}    <div class="item active">        
    {{else}}    <div class="item">{{end}}
          <a href="{{.Url}}"><img src="{{.Img}}" alt="Image"></a>
          <div class="carousel-caption">
            <a href="{{.Url}}"><h3>{{.Title}}</h3></a>
          </div>
        </div>
{{end}}{{end}}
      </div>
      {{end}}
`

const listtmpl = `
{{define "listnews"}}
<div class="col-sm-4">
{{range .}} {{if .Title}}
    <div class="well">
		<a href='{{.Url}}'>
		<p>{{.Title}}</p>
		</a>
    </div>
{{end}}{{end}}
</div>
{{end}}

`

const newstmpl = `
{{define "news"}}
<div class="container">
  <h4 class="channel-header">Apps</h4>
  <div class="row">
{{range .}} {{if .Title}}
    <div class="col-sm-4">
      <div class="panel panel-primary">
        <div class="panel-heading"><a href="{{.Url}}">{{.Title}}</a></div>
        <div class="panel-body"><a href="{{.Url}}"><img src="{{.Img}}" class="img-responsive" style="width:100%" alt="Image"></a></div>
        <div class="panel-footer">{{.Desc}}</div>
      </div>
    </div>
{{end}}{{end}}
  </div>
  <hr>
</div>
{{end}}
`
const newstmpl2 = `
{{define "news2"}}
<div class="container">
  <h4 class="channel-header">Gears</h4>
  <div class="row">
{{range .}} {{if .Title}}
    <div class="col-sm-4">
      <div class="panel panel-success">
        <div class="panel-heading"><a href="{{.Url}}">{{.Title}}</a></div>
        <div class="panel-body"><a href="{{.Url}}"><img src="{{.Img}}" class="img-responsive" style="width:100%" alt="Image"></a></div>
        <div class="panel-footer">{{.Desc}}</div>
      </div>
    </div>
{{end}}{{end}}
  </div>
  <hr>
</div>
{{end}}
`

const hometmpl = `
{{define "home"}}
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewpoint" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="http://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
<link rel="icon" type="image/png" href="https://camo.githubusercontent.com/bc8853eda7cbe8b23b35b6b607b15367af9c3c36/68747470733a2f2f7261772e6769746875622e636f6d2f676f6c616e672d73616d706c65732f676f706865722d766563746f722f6d61737465722f676f706865722d66726f6e742e706e67">
<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.12.4/jquery.min.js"></script>
<script src="http://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
<style>
	h4.channel-header{
		color:#9d9d9d;
	}
	.navbar-inverse {
	  background-color: #E0EBF5;	
	}
	a.navbar-brand img{
		height:50px;
	}
	.carousel-caption a, .panel-heading a{
		color:#fff;
	}
	.panel-success>.panel-heading a{
		color:#3c763d;
	}
	div.well a{
		color:#333;
	}	
    /* Add a gray background color and some padding to the footer */
    footer {
      background-color: #f2f2f2;
      padding: 25px;
    }
    .carousel-inner img {
      width: 100%; /* Set width to 100% */
      min-height: 200px;
    }
    /* Hide the carousel text when the screen is less than 600 pixels wide */
    @media (max-width: 600px) {
      .carousel-caption {
        display: none;
      }
    }
</style>
</head>
<body>
<nav class="navbar navbar-inverse">
  <div class="container-fluid">
    <div class="navbar-header">
      <button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#myNavbar">
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
      </button>
      <a class="navbar-brand" style="padding-top:0px" href="#"><img src="http://www.unixstickers.com/image/cache/data/stickers/golang/Go-white.sh-600x600.png"></a>
    </div>
    <div class="collapse navbar-collapse" id="myNavbar">
      <ul class="nav navbar-nav">
        <li class="active"><a href="#">Home</a></li>
        <li><a href="#">About</a></li>
        <li><a href="https://golang.org/">GoLang</a></li>
        <li><a href="http://ec2-52-207-252-146.compute-1.amazonaws.com/">Django</a></li>
      </ul>
      <ul class="nav navbar-nav navbar-right">
        <li><a href="#"><span class="glyphicon glyphicon-log-in"></span> Login</a></li>
      </ul>
    </div>
  </div>
</nav>
<div class="container">
<div class="row">
  <div class="col-sm-8">
    <div id="myCarousel" class="carousel slide" data-ride="carousel">
      <!-- Indicators -->
      <ol class="carousel-indicators">
        <li data-target="#myCarousel" data-slide-to="0" class="active"></li>
        <li data-target="#myCarousel" data-slide-to="1"></li>
        <li data-target="#myCarousel" data-slide-to="2"></li>
      </ol>

      <!-- Wrapper for slides -->
	{{template "carousel".Carousel}}
      <!-- Left and right controls -->
	  
      <a class="left carousel-control" href="#myCarousel" role="button" data-slide="prev">
        <span class="glyphicon glyphicon-chevron-left" aria-hidden="true"></span>
        <span class="sr-only">Previous</span>
      </a>
      <a class="right carousel-control" href="#myCarousel" role="button" data-slide="next">
        <span class="glyphicon glyphicon-chevron-right" aria-hidden="true"></span>
        <span class="sr-only">Next</span>
      </a>
    </div>
  </div>
	<!-- top right START-->
{{template "listnews".ListNews}}
	<!-- top right END-->  
</div>
<hr>
</div>

	<!-- news channel 1 START-->
{{template "news".Channel1}}
	<!-- news channel 1 END-->  

	<!-- news channel 2 START-->
{{template "news2".Channel2}}
	<!-- news channel 2 END-->  

<footer class="container-fluid text-center">
  <p>Created in Golang</p>
</footer>
</body>
</html>
{{end}}
`
