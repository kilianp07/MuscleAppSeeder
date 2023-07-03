package randomVideo

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

// GetRandomYouTubeVideos retrieves a specified number of random YouTube video links.
func Get(numVideos int) ([]string, error) {
	// Create a new YouTube service with the provided API key
	client := &http.Client{
		Transport: &transport.APIKey{Key: os.Getenv("YOUTUBE_API_KEY")},
	}
	service, err := youtube.New(client)
	if err != nil {
		return nil, fmt.Errorf("Unable to create YouTube service: %v", err)
	}

	// Create a new search call to search for random videos
	searchCall := service.Search.List([]string{"id"}).
		Q("Fitness").
		Type("video").
		MaxResults(int64(numVideos)) // Increase this number if you want more search results to choose from

	// Execute the search call
	response, err := searchCall.Do()
	if err != nil {
		return nil, fmt.Errorf("Error calling YouTube API: %v", err)
	}

	// Get the list of video IDs from the search response
	videoIDs := []string{}
	for _, item := range response.Items {
		videoIDs = append(videoIDs, item.Id.VideoId)
	}

	// Shuffle the video IDs
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(videoIDs), func(i, j int) {
		videoIDs[i], videoIDs[j] = videoIDs[j], videoIDs[i]
	})

	// Select the specified number of random video IDs
	selectedVideoIDs := []string{}
	if numVideos < len(videoIDs) {
		selectedVideoIDs = videoIDs[:numVideos]
	} else {
		selectedVideoIDs = videoIDs
	}

	// Generate the video links
	videoLinks := []string{}
	for _, videoID := range selectedVideoIDs {
		videoLinks = append(videoLinks, fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoID))
	}

	return videoLinks, nil
}
