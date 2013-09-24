import json
import requests
from BeautifulSoup import BeautifulSoup
import os

if __name__ == '__main__':
	
	if os.path.isfile('course_list.html'):
		with open('course_list.html', 'r') as f:
			html = f.read()
	else:
		r = requests.get("http://www.heinz.cmu.edu/academic-resources/course-results/index.aspx")

		if r.status_code != 200:
			print 'Something went wrong with the http request'
			sys.exit()

		html = r.content

	

	soup = BeautifulSoup(html)

	courses = []

	for line in soup.find(id="specialtbl").findAll("tr"):
			info = line.findAll("a")
			if len(info) >= 3:
				
				course = dict(
					num = info[0].getText(),
					title = info[1]['title'],
					professor = info[2].getText(),
					subject = line.findAll("td")[-1].getText()
				)

				courses.append(course)

	print courses[0]
	print len(courses)

	with open("classes.json", 'w') as f:
		classes_json = json.dumps(courses)
		f.write(classes_json)