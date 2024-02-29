import pandas as pd
import requests
pd.set_option('display.max_columns', None)
test_url = 'https://technicalowl.herokuapp.com/playerhustlestats?seasontype=regular&season=2023'
class Player:
    def __init__(self, player_id, player_name, gp, seconds_played, contested_shots_2pt, contested_shots_3pt,
                 deflections, charges_drawn, screen_ast_pts, loose_balls_recovered, box_outs, distance_traveled,
                 avg_speed, avg_speed_zscore, contested_shots_2pt_per_game_zscore, contested_shots_3pt_per_game_zscore,
                 deflections_per_game_zscore, charges_drawn_per_game_zscore, screen_ast_pts_per_game_zscore,
                 loose_balls_recovered_per_game_zscore, box_outs_per_game_zscore, distance_traveled_per_game_zscore,
                 contested_shots_2pt_total_zscore, contested_shots_3pt_total_zscore, deflections_total_zscore,
                 charges_drawn_total_zscore, screen_ast_pts_total_zscore, loose_balls_recovered_total_zscore,
                 box_outs_total_zscore, distance_traveled_total_zscore, contested_shots_2pt_per_36_zscore,
                 contested_shots_3pt_per_36_zscore, deflections_per_36_zscore, charges_drawn_per_36_zscore,
                 screen_ast_pts_per_36_zscore, loose_balls_recovered_per_36_zscore, box_outs_per_36_zscore,
                 distance_traveled_per_36_zscore):
        self.player_id = player_id
        self.player_name = player_name
        self.gp = gp
        self.seconds_played = seconds_played
        self.contested_shots_2pt = contested_shots_2pt
        self.contested_shots_3pt = contested_shots_3pt
        self.deflections = deflections
        self.charges_drawn = charges_drawn
        self.screen_ast_pts = screen_ast_pts
        self.loose_balls_recovered = loose_balls_recovered
        self.box_outs = box_outs
        self.distance_traveled = distance_traveled
        self.avg_speed = avg_speed
        self.avg_speed_zscore = avg_speed_zscore
        self.contested_shots_2pt_per_game_zscore = contested_shots_2pt_per_game_zscore
        self.contested_shots_3pt_per_game_zscore = contested_shots_3pt_per_game_zscore
        self.deflections_per_game_zscore = deflections_per_game_zscore
        self.charges_drawn_per_game_zscore = charges_drawn_per_game_zscore
        self.screen_ast_pts_per_game_zscore = screen_ast_pts_per_game_zscore
        self.loose_balls_recovered_per_game_zscore = loose_balls_recovered_per_game_zscore
        self.box_outs_per_game_zscore = box_outs_per_game_zscore
        self.distance_traveled_per_game_zscore = distance_traveled_per_game_zscore
        self.contested_shots_2pt_total_zscore = contested_shots_2pt_total_zscore
        self.contested_shots_3pt_total_zscore = contested_shots_3pt_total_zscore
        self.deflections_total_zscore = deflections_total_zscore
        self.charges_drawn_total_zscore = charges_drawn_total_zscore
        self.screen_ast_pts_total_zscore = screen_ast_pts_total_zscore
        self.loose_balls_recovered_total_zscore = loose_balls_recovered_total_zscore
        self.box_outs_total_zscore = box_outs_total_zscore
        self.distance_traveled_total_zscore = distance_traveled_total_zscore
        self.contested_shots_2pt_per_36_zscore = contested_shots_2pt_per_36_zscore
        self.contested_shots_3pt_per_36_zscore = contested_shots_3pt_per_36_zscore
        self.deflections_per_36_zscore = deflections_per_36_zscore
        self.charges_drawn_per_36_zscore = charges_drawn_per_36_zscore
        self.screen_ast_pts_per_36_zscore = screen_ast_pts_per_36_zscore
        self.loose_balls_recovered_per_36_zscore = loose_balls_recovered_per_36_zscore
        self.box_outs_per_36_zscore = box_outs_per_36_zscore
        self.distance_traveled_per_36_zscore = distance_traveled_per_36_zscore
def popular_objeto_com_json(r):

    # Extrai os dados do JSON
    player_id = r['player_id']
    player_name = r['player_name']
    gp = r['gp']
    seconds_played = r['seconds_played']
    contested_shots_2pt = r['contested_shots_2pt']
    contested_shots_3pt = r['contested_shots_3pt']
    deflections = r['deflections']
    charges_drawn = r['charges_drawn']
    screen_ast_pts = r['screen_ast_pts']
    loose_balls_recovered = r['loose_balls_recovered']
    box_outs = r['box_outs']
    distance_traveled = r['distance_traveled']
    avg_speed = r['avg_speed']
    avg_speed_zscore = r['avg_speed_zscore']
    contested_shots_2pt_per_game_zscore = r['contested_shots_2pt_per_game_zscore']
    contested_shots_3pt_per_game_zscore = r['contested_shots_3pt_per_game_zscore']
    deflections_per_game_zscore = r['deflections_per_game_zscore']
    charges_drawn_per_game_zscore = r['charges_drawn_per_game_zscore']
    screen_ast_pts_per_game_zscore = r['screen_ast_pts_per_game_zscore']
    loose_balls_recovered_per_game_zscore = r['loose_balls_recovered_per_game_zscore']
    box_outs_per_game_zscore = r['box_outs_per_game_zscore']
    distance_traveled_per_game_zscore = r['distance_traveled_per_game_zscore']
    contested_shots_2pt_total_zscore = r['contested_shots_2pt_total_zscore']
    contested_shots_3pt_total_zscore = r['contested_shots_3pt_total_zscore']
    deflections_total_zscore = r['deflections_total_zscore']
    charges_drawn_total_zscore = r['charges_drawn_total_zscore']
    screen_ast_pts_total_zscore = r['screen_ast_pts_total_zscore']
    loose_balls_recovered_total_zscore = r['loose_balls_recovered_total_zscore']
    box_outs_total_zscore = r['box_outs_total_zscore']
    distance_traveled_total_zscore = r['distance_traveled_total_zscore']
    contested_shots_2pt_per_36_zscore = r['contested_shots_2pt_per_36_zscore']
    contested_shots_3pt_per_36_zscore = r['contested_shots_3pt_per_36_zscore']
    deflections_per_36_zscore = r['deflections_per_36_zscore']
    charges_drawn_per_36_zscore = r['charges_drawn_per_36_zscore']
    screen_ast_pts_per_36_zscore = r['screen_ast_pts_per_36_zscore']
    loose_balls_recovered_per_36_zscore = r['loose_balls_recovered_per_36_zscore']
    box_outs_per_36_zscore = r['box_outs_per_36_zscore']
    distance_traveled_per_36_zscore = r['distance_traveled_per_36_zscore']

    player = Player(player_id, player_name, gp, seconds_played, contested_shots_2pt, contested_shots_3pt,
                 deflections, charges_drawn, screen_ast_pts, loose_balls_recovered, box_outs, distance_traveled,
                 avg_speed, avg_speed_zscore, contested_shots_2pt_per_game_zscore, contested_shots_3pt_per_game_zscore,
                 deflections_per_game_zscore, charges_drawn_per_game_zscore, screen_ast_pts_per_game_zscore,
                 loose_balls_recovered_per_game_zscore, box_outs_per_game_zscore, distance_traveled_per_game_zscore,
                 contested_shots_2pt_total_zscore, contested_shots_3pt_total_zscore, deflections_total_zscore,
                 charges_drawn_total_zscore, screen_ast_pts_total_zscore, loose_balls_recovered_total_zscore,
                 box_outs_total_zscore, distance_traveled_total_zscore, contested_shots_2pt_per_36_zscore,
                 contested_shots_3pt_per_36_zscore, deflections_per_36_zscore, charges_drawn_per_36_zscore,
                 screen_ast_pts_per_36_zscore, loose_balls_recovered_per_36_zscore, box_outs_per_36_zscore,
                 distance_traveled_per_36_zscore)

    return player

r = requests.get(url = test_url).json()
#print(r['data'][0])
#print(r)
player = popular_objeto_com_json(r['data'][0])
print(player.charges_drawn)