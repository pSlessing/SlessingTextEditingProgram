package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/nsf/termbox-go"
	"os"
	"path/filepath"
)

// Settings represents the configuration structure for termbox colors
type Settings struct {
	BGColor          termbox.Attribute `json:"bg_color"`
	FGColor          termbox.Attribute `json:"fg_color"`
	StatusBGColor    termbox.Attribute `json:"status_bg_color"`
	StatusFGColor    termbox.Attribute `json:"status_fg_color"`
	MsgBGColor       termbox.Attribute `json:"msg_bg_color"`
	MsgFGColor       termbox.Attribute `json:"msg_fg_color"`
	LineCountBGColor termbox.Attribute `json:"line_count_bg_color"`
	LineCountFGColor termbox.Attribute `json:"line_count_fg_color"`
}

// GetDefaultSettings returns the default configuration
func GetDefaultSettings() Settings {
	return Settings{
		BGColor:          termbox.ColorBlack,
		FGColor:          termbox.ColorWhite,
		StatusBGColor:    termbox.ColorWhite,
		StatusFGColor:    termbox.ColorBlack,
		MsgBGColor:       termbox.ColorWhite,
		MsgFGColor:       termbox.ColorBlack,
		LineCountBGColor: termbox.ColorWhite,
		LineCountFGColor: termbox.ColorCyan,
	}
}

// SaveSettings saves the current settings to a JSON file
func SaveSettings(settings Settings) error {
	// Get OS-specific config directory
	configDir, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get user config directory: %w", err)
	}
	configDir = filepath.Join(configDir, "termbox-editor")

	// Ensure the config directory exists
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	// Convert settings to JSON string
	jsonStr, err := SettingsToJSON(settings)
	if err != nil {
		return fmt.Errorf("failed to convert settings to JSON: %w", err)
	}

	// Write to file using similar pattern to WriteBufferToFile
	configPath := filepath.Join(configDir, "config.json")
	file, err := os.Create(configPath)
	if err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	_, err = writer.WriteString(jsonStr)
	if err != nil {
		return fmt.Errorf("failed to write settings to file: %w", err)
	}

	return nil
}

// LoadSettings loads settings from a JSON file, creating default config if file doesn't exist
func LoadSettings() (Settings, error) {
	// Get OS-specific config directory
	configDir, err := os.UserConfigDir()
	if err != nil {
		return Settings{}, fmt.Errorf("failed to get user config directory: %w", err)
	}
	configDir = filepath.Join(configDir, "termbox-editor")

	configPath := filepath.Join(configDir, "config.json")

	// Try to open the file, similar to OpenFile pattern
	file, err := os.Open(configPath)
	if err != nil {
		// File doesn't exist, create default config
		defaultSettings := GetDefaultSettings()
		if saveErr := SaveSettings(defaultSettings); saveErr != nil {
			return Settings{}, fmt.Errorf("failed to create default config: %w", saveErr)
		}
		return defaultSettings, nil
	}
	defer file.Close()

	// Read the file content
	var jsonContent string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		jsonContent += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		return Settings{}, fmt.Errorf("failed to read config file: %w", err)
	}

	// Convert JSON to settings
	settings, err := JSONToSettings(jsonContent)
	if err != nil {
		return Settings{}, fmt.Errorf("failed to parse settings JSON: %w", err)
	}

	return settings, nil
}

// SettingsToJSON converts a Settings struct to a JSON string
func SettingsToJSON(settings Settings) (string, error) {
	jsonData, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal settings to JSON: %w", err)
	}
	return string(jsonData), nil
}

// JSONToSettings converts a JSON string to a Settings struct
func JSONToSettings(jsonStr string) (Settings, error) {
	var settings Settings
	if err := json.Unmarshal([]byte(jsonStr), &settings); err != nil {
		return Settings{}, fmt.Errorf("failed to unmarshal JSON to settings: %w", err)
	}
	return settings, nil
}

// ApplySettings applies the loaded settings to the global color variables
func ApplySettings(settings Settings) {
	BGCOLOR = settings.BGColor
	FGCOLOR = settings.FGColor
	STATUSBGCOLOR = settings.StatusBGColor
	STATUSFGCOLOR = settings.StatusFGColor
	MSGBGCOLOR = settings.MsgBGColor
	MSGFGCOLOR = settings.MsgFGColor
	LINECOUNTBGCOLOR = settings.LineCountBGColor
	LINECOUNTFGCOLOR = settings.LineCountFGColor
}

// GetCurrentSettings creates a Settings struct from the current global variables
func GetCurrentSettings() Settings {
	return Settings{
		BGColor:          BGCOLOR,
		FGColor:          FGCOLOR,
		StatusBGColor:    STATUSBGCOLOR,
		StatusFGColor:    STATUSFGCOLOR,
		MsgBGColor:       MSGBGCOLOR,
		MsgFGColor:       MSGFGCOLOR,
		LineCountBGColor: LINECOUNTBGCOLOR,
		LineCountFGColor: LINECOUNTFGCOLOR,
	}
}

// Example usage:
/*
func main() {
	// Load settings on startup (creates default if doesn't exist)
	settings, err := LoadSettings()
	if err != nil {
		fmt.Printf("Error loading settings: %v\n", err)
		return
	}
	ApplySettings(settings)

	// Your application code here...

	// Save settings when modified
	currentSettings := GetCurrentSettings()
	err = SaveSettings(currentSettings)
	if err != nil {
		fmt.Printf("Error saving settings: %v\n", err)
	}
}
*/
