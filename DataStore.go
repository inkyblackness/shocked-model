package model

import (
	"github.com/inkyblackness/res/audio"
)

// FailureFunc is for failed queries.
type FailureFunc func()

// DataStore describes the necessary methods for querying and modifying model data.
type DataStore interface {
	// Projects queries the currently available projects.
	Projects(onSuccess func(projects []string), onFailure FailureFunc)
	// NewProject requests to create a new project with given ID.
	NewProject(projectID string, onSuccess func(), onFailure FailureFunc)

	// SaveProject requests to persist all pending modifications.
	SaveProject(projectID string)

	// Font queries a specific font.
	Font(projectID string, fontID int, onSuccess func(font *Font), onFailure FailureFunc)

	// Bitmap queries the data of a bitmap resource.
	Bitmap(projectID string, key ResourceKey,
		onSuccess func(ResourceKey, *RawBitmap), onFailure FailureFunc)
	// SetBitmap requests to set the data of a bitmap resource.
	SetBitmap(projectID string, key ResourceKey, bmp *RawBitmap,
		onSuccess func(ResourceKey, *RawBitmap), onFailure FailureFunc)

	// Text queries the data of a text resource.
	Text(projectID string, key ResourceKey,
		onSuccess func(ResourceKey, string), onFailure FailureFunc)
	// SetText requests to set the data of a text resource.
	SetText(projectID string, key ResourceKey, text string,
		onSuccess func(ResourceKey, string), onFailure FailureFunc)

	// Audio queries the data of an audio resource.
	Audio(projectID string, key ResourceKey,
		onSuccess func(ResourceKey, audio.SoundData), onFailure FailureFunc)
	// SetAudio requests to set the audio of a text resource.
	SetAudio(projectID string, key ResourceKey, data audio.SoundData,
		onSuccess func(ResourceKey), onFailure FailureFunc)

	// GameObjects queries the basic properties of all objects in the project.
	GameObjects(projectID string, onSuccess func(objects []GameObject), onFailure FailureFunc)
	// GameObjectIcon queries the icon bitmap of a game object.
	GameObjectIcon(projectID string, class, subclass, objType int, onSuccess func(bmp *RawBitmap), onFailure FailureFunc)
	// SetGameObject requests to update the properties of a game object.
	SetGameObject(projectID string, class, subclass, objType int, properties *GameObjectProperties,
		onSuccess func(properties *GameObjectProperties), onFailure FailureFunc)
	// GameObjectBitmap queries the data of a bitmap resource for a game object.
	GameObjectBitmap(projectID string, class, subclass, objType int, index int,
		onSuccess func(*RawBitmap), onFailure FailureFunc)
	// SetGameObjectBitmap requests to set the data of a bitmap resource for a game object.
	SetGameObjectBitmap(projectID string, class, subclass, objType int, index int, bmp *RawBitmap,
		onSuccess func(), onFailure FailureFunc)

	// ElectronicMessage queries the data of a specific electronic message.
	ElectronicMessage(projectID string, messageType ElectronicMessageType, id int,
		onSuccess func(message ElectronicMessage), onFailure FailureFunc)
	// SetElectronicMessage requests to update the properties of a specific electronic message.
	SetElectronicMessage(projectID string, messageType ElectronicMessageType, id int, message ElectronicMessage,
		onSuccess func(message ElectronicMessage), onFailure FailureFunc)
	// ElectronicMessageAudio queries the audio part of a specific electronic message.
	ElectronicMessageAudio(projectID string, messageType ElectronicMessageType, id int, language ResourceLanguage,
		onSuccess func(data audio.SoundData), onFailure FailureFunc)
	// SetElectronicMessageAudio requests to update the audio part of a specific electronic message.
	SetElectronicMessageAudio(projectID string, messageType ElectronicMessageType, id int, language ResourceLanguage, data audio.SoundData,
		onSuccess func(), onFailure FailureFunc)
	// RemoveElectronicMessage requests to remove the identified message from the resources.
	RemoveElectronicMessage(projectID string, messageType ElectronicMessageType, id int,
		onSuccess func(), onFailure FailureFunc)

	// Palette queries a palette.
	Palette(projectID string, paletteID string, onSuccess func(colors [256]Color), onFailure FailureFunc)
	// Levels queries all levels of a project.
	Levels(projectID string, archiveID string, onSuccess func(levels []Level), onFailure FailureFunc)

	// LevelProperties requests the basic properties of a level.
	LevelProperties(projectID string, archiveID string, levelID int, onSuccess func(properties LevelProperties), onFailure FailureFunc)
	// SetLevelProperties requests to update basic properties of a level.
	SetLevelProperties(projectID string, archiveID string, levelID int, properties LevelProperties,
		onSuccess func(properties LevelProperties), onFailure FailureFunc)

	// LevelTextures queries the texture IDs for a level.
	LevelTextures(projectID string, archiveID string, levelID int, onSuccess func(textureIDs []int), onFailure FailureFunc)
	// SetLevelTextures requests to set the list of textures for a level.
	SetLevelTextures(projectID string, archiveID string, levelID int, textureIDs []int, onSuccess func(textureIDs []int), onFailure FailureFunc)
	// LevelTextureAnimations queries the animation properties of texture groups.
	LevelTextureAnimations(projectID string, archiveID string, levelID int,
		onSuccess func(animations []TextureAnimation), onFailure FailureFunc)
	// SetLevelTextureAnimation requests to change the properties of an animated texture group.
	SetLevelTextureAnimation(projectID string, archiveID string, levelID int, animationGroup int, properties TextureAnimation,
		onSuccess func(animations []TextureAnimation), onFailure FailureFunc)

	// Textures queries all texture information of a project.
	Textures(projectID string, onSuccess func(textures []TextureProperties), onFailure FailureFunc)
	// SetTextureProperties requests to change properties of a single texture.
	SetTextureProperties(projectID string, textureID int, newProperties *TextureProperties,
		onSuccess func(properties *TextureProperties), onFailure FailureFunc)
	// TextureBitmap queries the texture bitmap of a texture.
	TextureBitmap(projectID string, textureID int, size string, onSuccess func(bmp *RawBitmap), onFailure FailureFunc)
	// SetTextureBitmap requests to update the bitmap of a texture.
	SetTextureBitmap(projectID string, textureID int, size string, bmp *RawBitmap,
		onSuccess func(bmp *RawBitmap), onFailure FailureFunc)

	// Tiles queries the complete tile map of a level.
	Tiles(projectID string, archiveID string, levelID int, onSuccess func(tiles Tiles), onFailure FailureFunc)

	// Tile requests the properties of a specific tile.
	Tile(projectID string, archiveID string, levelID int, x, y int,
		onSuccess func(properties TileProperties), onFailure FailureFunc)
	// SetTile requests to update properties of a specific tile.
	SetTile(projectID string, archiveID string, levelID int, x, y int, properties TileProperties,
		onSuccess func(properties TileProperties), onFailure FailureFunc)

	// LevelObjects requests all objects of the level.
	LevelObjects(projectID string, archiveID string, levelID int,
		onSuccess func(objects *LevelObjects), onFailure FailureFunc)
	// AddLevelObject requests to add a new object to the level.
	AddLevelObject(projectID string, archiveID string, levelID int, template LevelObjectTemplate,
		onSuccess func(object LevelObject), onFailure FailureFunc)
	// RemoveLevelObject requests to remove an object from the level.
	RemoveLevelObject(projectID string, archiveID string, levelID int, objectID int,
		onSuccess func(), onFailure FailureFunc)
	// SetLevelObject requests to update object properties.
	SetLevelObject(projectID string, archiveID string, levelID int, objectID int, properties *LevelObjectProperties,
		onSuccess func(properties *LevelObjectProperties), onFailure FailureFunc)

	// LevelSurveillanceObjects requests all surveillance objects of the identified level.
	LevelSurveillanceObjects(projectID string, archiveID string, levelID int,
		onSuccess func(objects []SurveillanceObject), onFailure FailureFunc)
	// SetLevelSurveillanceObject requests to set the properties of one surveillance object.
	SetLevelSurveillanceObject(projectID string, archiveID string, levelID int, surveillanceIndex int, data SurveillanceObject,
		onSuccess func(objects []SurveillanceObject), onFailure FailureFunc)
}
