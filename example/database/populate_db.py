import json
import sqlite3
import random

char_dir = json.loads(open("graph_nodes.json","rb").read())
db = sqlite3.connect("example.db")
cur = db.cursor()
cur.execute("delete from character;")
cur.execute("delete from episode;")
cur.execute("delete from appearance;")

for each in char_dir.keys():
    print "Character %s" % each
    cur.execute("insert into character (id, name) values (?,?);", (random.randint(1,10**10),each))
    for episode in char_dir[each]:
        cur.execute("select count(title) from episode where title=\"%s\";" % episode)
        repeats, = cur.__iter__().next()
        if repeats <1:
            cur.execute("insert into episode (id, title) values (?,?);", (random.randint(1,10**10),episode))
        """ adding appearance record """
        cur.execute("select id from episode where title=\"%s\";" % episode)
        episode_id, = cur.__iter__().next()
        cur.execute("select id from character where name=\"%s\";" % each)
        character_id, = cur.__iter__().next()
        # print 'episode id: ', episode_id
        # print 'character id: ', character_id
        cur.execute("insert into appearance (episode_id, character_id) values (?,?);", (episode_id, character_id))

    db.commit()

db.close()