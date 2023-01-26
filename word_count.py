# from english_words import english_words_set
import random

english_words_set = []
common_words = [
"the", "be", "to", "of", "and", "a", "in", "that", "have", "I",
"it", "for", "not", "on" ,"with", "he", "as", "you", "do", "at",
"this", "but", "his", "by", "from", "they", "we", "say", "her", "she",
"or", "an", "will", "my", "one", "all", "would", "there", "their", "what",
"so", "up", "out", "if", "about", "who", "get", "which", "go", "me"
"when", "make", "can", "like", "time", "no", "just", "him", "know", "take"
"people", "into", "year", "your", "good", "some", "could", "them", "see", "other"
"than", "then", "now", "look", "only", "come", "its", "over", "think", "also"
"back", "after", "use", "two", "how", "our", "work", "first", "well", "way"
"even", "new", "want", "because", "any", "these", "give", "day", "most", "us"
]

comment = input() #this is the comment that will be evaluated 
words_list = comment.split()
word_count = 0
char_count = 0
english_word_count = 0
common_word_count = 0
word_char_count_list = []
word_count = len(comment.split())
char_count = len(comment)
base_points = 100
raw_points = base_points + word_count + char_count
extra_points = 0

for word in words_list:
	if word in english_words_set:
		extra_points += 1
		english_word_count += 1
		if word in common_words:
			extra_points -= 1
			common_word_count += 1
		
	word_char_count_list.append(len(word))


word_char_total = sum(word_char_count_list)
word_char_avg = word_char_total/word_count
total_submission_points = raw_points + extra_points

print("\n# of characters: ", end = "")
print(char_count)
print("\ntotal # of words: ", end ="")
print(word_count)
print("\n# of words found in the English dictionary [awarded 1 extra point]: ", end ="")
print(english_word_count)
print("\n# of Common words (top 100) [not awarded extra points]: ", end ="")
print(common_word_count)
print("\navg # of characters per word: ", end = "")
print(word_char_avg)
print("\nlength of words: ", end="")
print(word_char_count_list, "\n")
print("\raw points: ", end = "")
print(raw_points)
print("\nextra points: ", end = "")
print(extra_points, "\n")
print("\ntotal submission points: ", end = "")
print(total_submission_points, "\n")


votes = ["valid", "invalid"]
prompt = ["if", "and", "but", ""]
valid_count = 0
invalid_count = 0
confidence = 0
confidence_list = []
valid_points = 0
invalid_points = 0
raw_votes = input("\n# of votes to randomize and calculate?: ")
int_votes = int(raw_votes)


while (int_votes > 0):
	temp_vote = random.randint(0, 1)
	confidence = random.randint(0, 100)
	confidence_list.append(confidence)
	if votes[temp_vote] == votes[0]:
		valid_count += 1
		valid_points += (50+(1*(confidence/100)))
	else:
		invalid_count += 1
		invalid_points += (50+(1*(confidence/100)))
	int_votes -= 1
avg_confidence = sum(confidence_list)/len(confidence_list)

print("\n\nAverage confidence: ", end="")
print(avg_confidence)
print("\n# Valid votes: ", end="")
print(valid_count)
print("# Valid points: ", end="")
print(valid_points)
print("\n# Invalid votes: ", end="")
print(invalid_count)
print("# Invalid points: ", end="")
print(invalid_points)
total_engagement_points = int(valid_points) + int(invalid_points)
total_points = total_submission_points + total_engagement_points 
print("\n\n# Total points: ", end="")
print(total_points)