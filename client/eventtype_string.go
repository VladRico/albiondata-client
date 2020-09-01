// Code generated by "stringer -type=EventType"; DO NOT EDIT.

package client

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[evUnused-0]
	_ = x[evLeave-1]
	_ = x[evJoinFinished-2]
	_ = x[evMove-3]
	_ = x[evTeleport-4]
	_ = x[evChangeEquipment-5]
	_ = x[evHealthUpdate-6]
	_ = x[evEnergyUpdate-7]
	_ = x[evDamageShieldUpdate-8]
	_ = x[evCraftingFocusUpdate-9]
	_ = x[evActiveSpellEffectsUpdate-10]
	_ = x[evResetCooldowns-11]
	_ = x[evAttack-12]
	_ = x[evCastStart-13]
	_ = x[evChannelingUpdate-14]
	_ = x[evCastCancel-15]
	_ = x[evCastTimeUpdate-16]
	_ = x[evCastFinished-17]
	_ = x[evCastSpell-18]
	_ = x[evCastHit-19]
	_ = x[evCastHits-20]
	_ = x[evChannelingEnded-21]
	_ = x[evAttackBuilding-22]
	_ = x[evInventoryPutItem-23]
	_ = x[evInventoryDeleteItem-24]
	_ = x[evNewCharacter-25]
	_ = x[evNewEquipmentItem-26]
	_ = x[evNewSimpleItem-27]
	_ = x[evNewFurnitureItem-28]
	_ = x[evNewJournalItem-29]
	_ = x[evNewLaborerItem-30]
	_ = x[evNewSimpleHarvestableObject-31]
	_ = x[evNewSimpleHarvestableObjectList-32]
	_ = x[evNewHarvestableObject-33]
	_ = x[evNewSilverObject-34]
	_ = x[evNewBuilding-35]
	_ = x[evHarvestableChangeState-36]
	_ = x[evMobChangeState-37]
	_ = x[evFactionBuildingInfo-38]
	_ = x[evCraftBuildingInfo-39]
	_ = x[evRepairBuildingInfo-40]
	_ = x[evMeldBuildingInfo-41]
	_ = x[evConstructionSiteInfo-42]
	_ = x[evPlayerBuildingInfo-43]
	_ = x[evFarmBuildingInfo-44]
	_ = x[evTutorialBuildingInfo-45]
	_ = x[evLaborerObjectInfo-46]
	_ = x[evLaborerObjectJobInfo-47]
	_ = x[evMarketPlaceBuildingInfo-48]
	_ = x[evHarvestStart-49]
	_ = x[evHarvestCancel-50]
	_ = x[evHarvestFinished-51]
	_ = x[evTakeSilver-52]
	_ = x[evActionOnBuildingStart-53]
	_ = x[evActionOnBuildingCancel-54]
	_ = x[evActionOnBuildingFinished-55]
	_ = x[evItemRerollQualityStart-56]
	_ = x[evItemRerollQualityCancel-57]
	_ = x[evItemRerollQualityFinished-58]
	_ = x[evInstallResourceStart-59]
	_ = x[evInstallResourceCancel-60]
	_ = x[evInstallResourceFinished-61]
	_ = x[evCraftItemFinished-62]
	_ = x[evLogoutCancel-63]
	_ = x[evChatMessage-64]
	_ = x[evChatSay-65]
	_ = x[evChatWhisper-66]
	_ = x[evChatMuted-67]
	_ = x[evPlayEmote-68]
	_ = x[evStopEmote-69]
	_ = x[evSystemMessage-70]
	_ = x[evUtilityTextMessage-71]
	_ = x[evUpdateMoney-72]
	_ = x[evUpdateFame-73]
	_ = x[evUpdateLearningPoints-74]
	_ = x[evUpdateReSpecPoints-75]
	_ = x[evUpdateCurrency-76]
	_ = x[evUpdateFactionStanding-77]
	_ = x[evRespawn-78]
	_ = x[evServerDebugLog-79]
	_ = x[evCharacterEquipmentChanged-80]
	_ = x[evRegenerationHealthChanged-81]
	_ = x[evRegenerationEnergyChanged-82]
	_ = x[evRegenerationMountHealthChanged-83]
	_ = x[evRegenerationCraftingChanged-84]
	_ = x[evRegenerationHealthEnergyComboChanged-85]
	_ = x[evRegenerationPlayerComboChanged-86]
	_ = x[evDurabilityChanged-87]
	_ = x[evNewLoot-88]
	_ = x[evAttachItemContainer-89]
	_ = x[evDetachItemContainer-90]
	_ = x[evInvalidateItemContainer-91]
	_ = x[evLockItemContainer-92]
	_ = x[evGuildUpdate-93]
	_ = x[evGuildPlayerUpdated-94]
	_ = x[evInvitedToGuild-95]
	_ = x[evGuildMemberWorldUpdate-96]
	_ = x[evUpdateMatchDetails-97]
	_ = x[evObjectEvent-98]
	_ = x[evNewMonolithObject-99]
	_ = x[evNewSiegeCampObject-100]
	_ = x[evNewOrbObject-101]
	_ = x[evNewCastleObject-102]
	_ = x[evNewSpellEffectArea-103]
	_ = x[evUpdateSpellEffectArea-104]
	_ = x[evNewChainSpell-105]
	_ = x[evUpdateChainSpell-106]
	_ = x[evNewTreasureChest-107]
	_ = x[evStartMatch-108]
	_ = x[evStartTerritoryMatchInfos-109]
	_ = x[evStartArenaMatchInfos-110]
	_ = x[evEndTerritoryMatch-111]
	_ = x[evEndArenaMatch-112]
	_ = x[evMatchUpdate-113]
	_ = x[evActiveMatchUpdate-114]
	_ = x[evNewMob-115]
	_ = x[evDebugAggroInfo-116]
	_ = x[evDebugVariablesInfo-117]
	_ = x[evDebugReputationInfo-118]
	_ = x[evDebugDiminishingReturnInfo-119]
	_ = x[evDebugSmartClusterQueueInfo-120]
	_ = x[evClaimOrbStart-121]
	_ = x[evClaimOrbFinished-122]
	_ = x[evClaimOrbCancel-123]
	_ = x[evOrbUpdate-124]
	_ = x[evOrbClaimed-125]
	_ = x[evNewWarCampObject-126]
	_ = x[evNewMatchLootChestObject-127]
	_ = x[evNewArenaExit-128]
	_ = x[evGuildMemberTerritoryUpdate-129]
	_ = x[evInvitedMercenaryToMatch-130]
	_ = x[evClusterInfoUpdate-131]
	_ = x[evForcedMovement-132]
	_ = x[evForcedMovementCancel-133]
	_ = x[evCharacterStats-134]
	_ = x[evCharacterStatsKillHistory-135]
	_ = x[evCharacterStatsDeathHistory-136]
	_ = x[evGuildStats-137]
	_ = x[evKillHistoryDetails-138]
	_ = x[evFullAchievementInfo-139]
	_ = x[evFinishedAchievement-140]
	_ = x[evAchievementProgressInfo-141]
	_ = x[evFullAchievementProgressInfo-142]
	_ = x[evFullTrackedAchievementInfo-143]
	_ = x[evFullAutoLearnAchievementInfo-144]
	_ = x[evQuestGiverQuestOffered-145]
	_ = x[evQuestGiverDebugInfo-146]
	_ = x[evConsoleEvent-147]
	_ = x[evTimeSync-148]
	_ = x[evChangeAvatar-149]
	_ = x[evChangeMountSkin-150]
	_ = x[evGameEvent-151]
	_ = x[evKilledPlayer-152]
	_ = x[evDied-153]
	_ = x[evKnockedDown-154]
	_ = x[evMatchPlayerJoinedEvent-155]
	_ = x[evMatchPlayerStatsEvent-156]
	_ = x[evMatchPlayerStatsCompleteEvent-157]
	_ = x[evMatchTimeLineEventEvent-158]
	_ = x[evMatchPlayerMainGearStatsEvent-159]
	_ = x[evMatchPlayerChangedAvatarEvent-160]
	_ = x[evInvitationPlayerTrade-161]
	_ = x[evPlayerTradeStart-162]
	_ = x[evPlayerTradeCancel-163]
	_ = x[evPlayerTradeUpdate-164]
	_ = x[evPlayerTradeFinished-165]
	_ = x[evPlayerTradeAcceptChange-166]
	_ = x[evMiniMapPing-167]
	_ = x[evMarketPlaceNotification-168]
	_ = x[evDuellingChallengePlayer-169]
	_ = x[evNewDuellingPost-170]
	_ = x[evDuelStarted-171]
	_ = x[evDuelEnded-172]
	_ = x[evDuelDenied-173]
	_ = x[evDuelLeftArea-174]
	_ = x[evDuelReEnteredArea-175]
	_ = x[evNewRealEstate-176]
	_ = x[evMiniMapOwnedBuildingsPositions-177]
	_ = x[evRealEstateListUpdate-178]
	_ = x[evGuildLogoUpdate-179]
	_ = x[evGuildLogoChanged-180]
	_ = x[evPlaceableObjectPlace-181]
	_ = x[evPlaceableObjectPlaceCancel-182]
	_ = x[evFurnitureObjectBuffProviderInfo-183]
	_ = x[evFurnitureObjectCheatProviderInfo-184]
	_ = x[evFarmableObjectInfo-185]
	_ = x[evNewUnreadMails-186]
	_ = x[evUnknown187-187]
	_ = x[evGuildLogoObjectUpdate-188]
	_ = x[evStartLogout-189]
	_ = x[evNewChatChannels-190]
	_ = x[evJoinedChatChannel-191]
	_ = x[evLeftChatChannel-192]
	_ = x[evRemovedChatChannel-193]
	_ = x[evAccessStatus-194]
	_ = x[evMounted-195]
	_ = x[evMountStart-196]
	_ = x[evMountCancel-197]
	_ = x[evNewTravelpoint-198]
	_ = x[evNewIslandAccessPoint-199]
	_ = x[evNewExit-200]
	_ = x[evUpdateHome-201]
	_ = x[evUpdateChatSettings-202]
	_ = x[evResurrectionOffer-203]
	_ = x[evResurrectionReply-204]
	_ = x[evLootEquipmentChanged-205]
	_ = x[evUpdateUnlockedGuildLogos-206]
	_ = x[evUpdateUnlockedAvatars-207]
	_ = x[evUpdateUnlockedAvatarRings-208]
	_ = x[evUpdateUnlockedBuildings-209]
	_ = x[evNewIslandManagement-210]
	_ = x[evNewTeleportStone-211]
	_ = x[evCloak-212]
	_ = x[evPartyInvitation-213]
	_ = x[evPartyJoined-214]
	_ = x[evPartyDisbanded-215]
	_ = x[evPartyPlayerJoined-216]
	_ = x[evPartyChangedOrder-217]
	_ = x[evPartyPlayerLeft-218]
	_ = x[evPartyLeaderChanged-219]
	_ = x[evPartyLootSettingChangedPlayer-220]
	_ = x[evPartySilverGained-221]
	_ = x[evPartyPlayerUpdated-222]
	_ = x[evPartyInvitationPlayerBusy-223]
	_ = x[evPartyMarkedObjectsUpdated-224]
	_ = x[evPartyOnClusterPartyJoined-225]
	_ = x[evPartySetRoleFlag-226]
	_ = x[evSpellCooldownUpdate-227]
	_ = x[evNewHellgate-228]
	_ = x[evNewHellgateExit-229]
	_ = x[evNewExpeditionExit-230]
	_ = x[evNewExpeditionNarrator-231]
	_ = x[evExitEnterStart-232]
	_ = x[evExitEnterCancel-233]
	_ = x[evExitEnterFinished-234]
	_ = x[evHellClusterTimeUpdate-235]
	_ = x[evNewQuestGiverObject-236]
	_ = x[evFullQuestInfo-237]
	_ = x[evQuestProgressInfo-238]
	_ = x[evQuestGiverInfoForPlayer-239]
	_ = x[evFullExpeditionInfo-240]
	_ = x[evExpeditionQuestProgressInfo-241]
	_ = x[evInvitedToExpedition-242]
	_ = x[evExpeditionRegistrationInfo-243]
	_ = x[evEnteringExpeditionStart-244]
	_ = x[evEnteringExpeditionCancel-245]
	_ = x[evRewardGranted-246]
	_ = x[evArenaRegistrationInfo-247]
	_ = x[evEnteringArenaStart-248]
	_ = x[evEnteringArenaCancel-249]
	_ = x[evEnteringArenaLockStart-250]
	_ = x[evEnteringArenaLockCancel-251]
	_ = x[evInvitedToArenaMatch-252]
	_ = x[evPlayerCounts-253]
	_ = x[evInCombatStateUpdate-254]
	_ = x[evOtherGrabbedLoot-255]
	_ = x[evSiegeCampClaimStart-256]
	_ = x[evSiegeCampClaimCancel-257]
	_ = x[evSiegeCampClaimFinished-258]
	_ = x[evSiegeCampScheduleResult-259]
	_ = x[evTreasureChestUsingStart-260]
	_ = x[evTreasureChestUsingFinished-261]
	_ = x[evTreasureChestUsingCancel-262]
	_ = x[evTreasureChestUsingOpeningComplete-263]
	_ = x[evTreasureChestForceCloseInventory-264]
	_ = x[evPremiumChanged-265]
	_ = x[evPremiumExtended-266]
	_ = x[evPremiumLifeTimeRewardGained-267]
	_ = x[evLaborerGotUpgraded-268]
	_ = x[evJournalGotFull-269]
	_ = x[evJournalFillError-270]
	_ = x[evFriendRequest-271]
	_ = x[evFriendRequestInfos-272]
	_ = x[evFriendInfos-273]
	_ = x[evFriendRequestAnswered-274]
	_ = x[evFriendOnlineStatus-275]
	_ = x[evFriendRequestCanceled-276]
	_ = x[evFriendRemoved-277]
	_ = x[evFriendUpdated-278]
	_ = x[evPartyLootItems-279]
	_ = x[evPartyLootItemsRemoved-280]
	_ = x[evReputationUpdate-281]
	_ = x[evDefenseUnitAttackBegin-282]
	_ = x[evDefenseUnitAttackEnd-283]
	_ = x[evDefenseUnitAttackDamage-284]
	_ = x[evUnrestrictedPvpZoneUpdate-285]
	_ = x[evReputationImplicationUpdate-286]
	_ = x[evNewMountObject-287]
	_ = x[evMountHealthUpdate-288]
	_ = x[evMountCooldownUpdate-289]
	_ = x[evNewExpeditionAgent-290]
	_ = x[evNewExpeditionCheckPoint-291]
	_ = x[evExpeditionStartEvent-292]
	_ = x[evVoteEvent-293]
	_ = x[evRatingEvent-294]
	_ = x[evNewArenaAgent-295]
	_ = x[evBoostFarmable-296]
	_ = x[evUseFunction-297]
	_ = x[evNewPortalEntrance-298]
	_ = x[evNewPortalExit-299]
	_ = x[evNewRandomDungeonExit-300]
	_ = x[evWaitingQueueUpdate-301]
	_ = x[evPlayerMovementRateUpdate-302]
	_ = x[evObserveStart-303]
	_ = x[evMinimapZergs-304]
	_ = x[evMinimapSmartClusterZergs-305]
	_ = x[evPaymentTransactions-306]
	_ = x[evPerformanceStatsUpdate-307]
	_ = x[evOverloadModeUpdate-308]
	_ = x[evDebugDrawEvent-309]
	_ = x[evRecordCameraMove-310]
	_ = x[evRecordStart-311]
	_ = x[evTerritoryClaimStart-312]
	_ = x[evTerritoryClaimCancel-313]
	_ = x[evTerritoryClaimFinished-314]
	_ = x[evTerritoryScheduleResult-315]
	_ = x[evUpdateAccountState-316]
	_ = x[evStartDeterministicRoam-317]
	_ = x[evGuildFullAccessTagsUpdated-318]
	_ = x[evGuildAccessTagUpdated-319]
	_ = x[evGvgSeasonUpdate-320]
	_ = x[evGvgSeasonCheatCommand-321]
	_ = x[evSeasonPointsByKillingBooster-322]
	_ = x[evFishingStart-323]
	_ = x[evFishingCast-324]
	_ = x[evFishingCatch-325]
	_ = x[evFishingFinished-326]
	_ = x[evFishingCancel-327]
	_ = x[evNewFloatObject-328]
	_ = x[evNewFishingZoneObject-329]
	_ = x[evFishingMiniGame-330]
	_ = x[evSteamAchievementCompleted-331]
	_ = x[evUpdatePuppet-332]
	_ = x[evChangeFlaggingFinished-333]
	_ = x[evNewOutpostObject-334]
	_ = x[evOutpostUpdate-335]
	_ = x[evOutpostClaimed-336]
	_ = x[evOutpostReward-337]
	_ = x[evOverChargeEnd-338]
	_ = x[evOverChargeStatus-339]
	_ = x[evPartyFinderFullUpdate-340]
	_ = x[evPartyFinderUpdate-341]
	_ = x[evPartyFinderApplicantsUpdate-342]
	_ = x[evPartyFinderEquipmentSnapshot-343]
	_ = x[evPartyFinderJoinRequestDeclined-344]
	_ = x[evNewUnlockedPersonalSeasonRewards-345]
	_ = x[evPersonalSeasonPointsGained-346]
	_ = x[evEasyAntiCheatMessageToClient-347]
	_ = x[evMatchLootChestOpeningStart-348]
	_ = x[evMatchLootChestOpeningFinished-349]
	_ = x[evMatchLootChestOpeningCancel-350]
	_ = x[evNotifyCrystalMatchReward-351]
	_ = x[evCrystalRealmFeedback-352]
	_ = x[evNewLocationMarker-353]
	_ = x[evNewTutorialBlocker-354]
	_ = x[evNewTileSwitch-355]
	_ = x[evNewInformationProvider-356]
	_ = x[evNewDynamicGuildLogo-357]
	_ = x[evTutorialUpdate-358]
	_ = x[evTriggerHintBox-359]
	_ = x[evRandomDungeonPositionInfo-360]
	_ = x[evNewLootChest-361]
	_ = x[evUpdateLootChest-362]
	_ = x[evLootChestOpened-363]
	_ = x[evNewShrine-364]
	_ = x[evUpdateShrine-365]
	_ = x[evMutePlayerUpdate-366]
	_ = x[evShopTileUpdate-367]
	_ = x[evShopUpdate-368]
	_ = x[evEasyAntiCheatKick-369]
	_ = x[evUnlockVanityUnlock-370]
	_ = x[evAvatarUnlocked-371]
	_ = x[evCustomizationChanged-372]
	_ = x[evBaseVaultInfo-373]
	_ = x[evGuildVaultInfo-374]
	_ = x[evBankVaultInfo-375]
	_ = x[evRecoveryVaultPlayerInfo-376]
	_ = x[evRecoveryVaultGuildInfo-377]
	_ = x[evUpdateWardrobe-378]
	_ = x[evCastlePhaseChanged-379]
	_ = x[evGuildAccountLogEvent-380]
	_ = x[evNewHideoutObject-381]
	_ = x[evNewHideoutManagement-382]
	_ = x[evNewHideoutExit-383]
	_ = x[evInitHideoutAttackStart-384]
	_ = x[evInitHideoutAttackCancel-385]
	_ = x[evInitHideoutAttackFinished-386]
	_ = x[evHideoutManagementUpdate-387]
	_ = x[evIpChanged-388]
	_ = x[evSmartClusterQueueUpdateInfo-389]
	_ = x[evSmartClusterQueueActiveInfo-390]
	_ = x[evSmartClusterQueueKickWarning-391]
	_ = x[evSmartClusterQueueInvite-392]
	_ = x[evReceivedGvgSeasonPoints-393]
	_ = x[evTerritoryBonusLevelUpdate-394]
	_ = x[evOpenWorldAttackScheduleStart-395]
	_ = x[evOpenWorldAttackScheduleFinished-396]
	_ = x[evOpenWorldAttackScheduleCancel-397]
	_ = x[evOpenWorldAttackConquerStart-398]
	_ = x[evOpenWorldAttackConquerFinished-399]
	_ = x[evOpenWorldAttackConquerCancel-400]
	_ = x[evOpenWorldAttackConquerStatus-401]
	_ = x[evOpenWorldAttackStart-402]
	_ = x[evOpenWorldAttackEnd-403]
	_ = x[evNewRandomResourceBlocker-404]
	_ = x[evNewHomeObject-405]
	_ = x[evHideoutObjectUpdate-406]
	_ = x[evUpdateInfamy-407]
	_ = x[evUnknown408-408]
	_ = x[evUnknown409-409]
	_ = x[evUnknown410-410]
	_ = x[evUnknown411-411]
	_ = x[evUnknown412-412]
	_ = x[evUnknown413-413]
	_ = x[evUnknown414-414]
	_ = x[evUnknown415-415]
	_ = x[evUnknown416-416]
	_ = x[evUnknown417-417]
	_ = x[evUnknown418-418]
	_ = x[evUnknown419-419]
	_ = x[evUnknown420-420]
	_ = x[evUnknown421-421]
	_ = x[evUnknown422-422]
}

const _EventType_name = "evUnusedevLeaveevJoinFinishedevMoveevTeleportevChangeEquipmentevHealthUpdateevEnergyUpdateevDamageShieldUpdateevCraftingFocusUpdateevActiveSpellEffectsUpdateevResetCooldownsevAttackevCastStartevChannelingUpdateevCastCancelevCastTimeUpdateevCastFinishedevCastSpellevCastHitevCastHitsevChannelingEndedevAttackBuildingevInventoryPutItemevInventoryDeleteItemevNewCharacterevNewEquipmentItemevNewSimpleItemevNewFurnitureItemevNewJournalItemevNewLaborerItemevNewSimpleHarvestableObjectevNewSimpleHarvestableObjectListevNewHarvestableObjectevNewSilverObjectevNewBuildingevHarvestableChangeStateevMobChangeStateevFactionBuildingInfoevCraftBuildingInfoevRepairBuildingInfoevMeldBuildingInfoevConstructionSiteInfoevPlayerBuildingInfoevFarmBuildingInfoevTutorialBuildingInfoevLaborerObjectInfoevLaborerObjectJobInfoevMarketPlaceBuildingInfoevHarvestStartevHarvestCancelevHarvestFinishedevTakeSilverevActionOnBuildingStartevActionOnBuildingCancelevActionOnBuildingFinishedevItemRerollQualityStartevItemRerollQualityCancelevItemRerollQualityFinishedevInstallResourceStartevInstallResourceCancelevInstallResourceFinishedevCraftItemFinishedevLogoutCancelevChatMessageevChatSayevChatWhisperevChatMutedevPlayEmoteevStopEmoteevSystemMessageevUtilityTextMessageevUpdateMoneyevUpdateFameevUpdateLearningPointsevUpdateReSpecPointsevUpdateCurrencyevUpdateFactionStandingevRespawnevServerDebugLogevCharacterEquipmentChangedevRegenerationHealthChangedevRegenerationEnergyChangedevRegenerationMountHealthChangedevRegenerationCraftingChangedevRegenerationHealthEnergyComboChangedevRegenerationPlayerComboChangedevDurabilityChangedevNewLootevAttachItemContainerevDetachItemContainerevInvalidateItemContainerevLockItemContainerevGuildUpdateevGuildPlayerUpdatedevInvitedToGuildevGuildMemberWorldUpdateevUpdateMatchDetailsevObjectEventevNewMonolithObjectevNewSiegeCampObjectevNewOrbObjectevNewCastleObjectevNewSpellEffectAreaevUpdateSpellEffectAreaevNewChainSpellevUpdateChainSpellevNewTreasureChestevStartMatchevStartTerritoryMatchInfosevStartArenaMatchInfosevEndTerritoryMatchevEndArenaMatchevMatchUpdateevActiveMatchUpdateevNewMobevDebugAggroInfoevDebugVariablesInfoevDebugReputationInfoevDebugDiminishingReturnInfoevDebugSmartClusterQueueInfoevClaimOrbStartevClaimOrbFinishedevClaimOrbCancelevOrbUpdateevOrbClaimedevNewWarCampObjectevNewMatchLootChestObjectevNewArenaExitevGuildMemberTerritoryUpdateevInvitedMercenaryToMatchevClusterInfoUpdateevForcedMovementevForcedMovementCancelevCharacterStatsevCharacterStatsKillHistoryevCharacterStatsDeathHistoryevGuildStatsevKillHistoryDetailsevFullAchievementInfoevFinishedAchievementevAchievementProgressInfoevFullAchievementProgressInfoevFullTrackedAchievementInfoevFullAutoLearnAchievementInfoevQuestGiverQuestOfferedevQuestGiverDebugInfoevConsoleEventevTimeSyncevChangeAvatarevChangeMountSkinevGameEventevKilledPlayerevDiedevKnockedDownevMatchPlayerJoinedEventevMatchPlayerStatsEventevMatchPlayerStatsCompleteEventevMatchTimeLineEventEventevMatchPlayerMainGearStatsEventevMatchPlayerChangedAvatarEventevInvitationPlayerTradeevPlayerTradeStartevPlayerTradeCancelevPlayerTradeUpdateevPlayerTradeFinishedevPlayerTradeAcceptChangeevMiniMapPingevMarketPlaceNotificationevDuellingChallengePlayerevNewDuellingPostevDuelStartedevDuelEndedevDuelDeniedevDuelLeftAreaevDuelReEnteredAreaevNewRealEstateevMiniMapOwnedBuildingsPositionsevRealEstateListUpdateevGuildLogoUpdateevGuildLogoChangedevPlaceableObjectPlaceevPlaceableObjectPlaceCancelevFurnitureObjectBuffProviderInfoevFurnitureObjectCheatProviderInfoevFarmableObjectInfoevNewUnreadMailsevUnknown187evGuildLogoObjectUpdateevStartLogoutevNewChatChannelsevJoinedChatChannelevLeftChatChannelevRemovedChatChannelevAccessStatusevMountedevMountStartevMountCancelevNewTravelpointevNewIslandAccessPointevNewExitevUpdateHomeevUpdateChatSettingsevResurrectionOfferevResurrectionReplyevLootEquipmentChangedevUpdateUnlockedGuildLogosevUpdateUnlockedAvatarsevUpdateUnlockedAvatarRingsevUpdateUnlockedBuildingsevNewIslandManagementevNewTeleportStoneevCloakevPartyInvitationevPartyJoinedevPartyDisbandedevPartyPlayerJoinedevPartyChangedOrderevPartyPlayerLeftevPartyLeaderChangedevPartyLootSettingChangedPlayerevPartySilverGainedevPartyPlayerUpdatedevPartyInvitationPlayerBusyevPartyMarkedObjectsUpdatedevPartyOnClusterPartyJoinedevPartySetRoleFlagevSpellCooldownUpdateevNewHellgateevNewHellgateExitevNewExpeditionExitevNewExpeditionNarratorevExitEnterStartevExitEnterCancelevExitEnterFinishedevHellClusterTimeUpdateevNewQuestGiverObjectevFullQuestInfoevQuestProgressInfoevQuestGiverInfoForPlayerevFullExpeditionInfoevExpeditionQuestProgressInfoevInvitedToExpeditionevExpeditionRegistrationInfoevEnteringExpeditionStartevEnteringExpeditionCancelevRewardGrantedevArenaRegistrationInfoevEnteringArenaStartevEnteringArenaCancelevEnteringArenaLockStartevEnteringArenaLockCancelevInvitedToArenaMatchevPlayerCountsevInCombatStateUpdateevOtherGrabbedLootevSiegeCampClaimStartevSiegeCampClaimCancelevSiegeCampClaimFinishedevSiegeCampScheduleResultevTreasureChestUsingStartevTreasureChestUsingFinishedevTreasureChestUsingCancelevTreasureChestUsingOpeningCompleteevTreasureChestForceCloseInventoryevPremiumChangedevPremiumExtendedevPremiumLifeTimeRewardGainedevLaborerGotUpgradedevJournalGotFullevJournalFillErrorevFriendRequestevFriendRequestInfosevFriendInfosevFriendRequestAnsweredevFriendOnlineStatusevFriendRequestCanceledevFriendRemovedevFriendUpdatedevPartyLootItemsevPartyLootItemsRemovedevReputationUpdateevDefenseUnitAttackBeginevDefenseUnitAttackEndevDefenseUnitAttackDamageevUnrestrictedPvpZoneUpdateevReputationImplicationUpdateevNewMountObjectevMountHealthUpdateevMountCooldownUpdateevNewExpeditionAgentevNewExpeditionCheckPointevExpeditionStartEventevVoteEventevRatingEventevNewArenaAgentevBoostFarmableevUseFunctionevNewPortalEntranceevNewPortalExitevNewRandomDungeonExitevWaitingQueueUpdateevPlayerMovementRateUpdateevObserveStartevMinimapZergsevMinimapSmartClusterZergsevPaymentTransactionsevPerformanceStatsUpdateevOverloadModeUpdateevDebugDrawEventevRecordCameraMoveevRecordStartevTerritoryClaimStartevTerritoryClaimCancelevTerritoryClaimFinishedevTerritoryScheduleResultevUpdateAccountStateevStartDeterministicRoamevGuildFullAccessTagsUpdatedevGuildAccessTagUpdatedevGvgSeasonUpdateevGvgSeasonCheatCommandevSeasonPointsByKillingBoosterevFishingStartevFishingCastevFishingCatchevFishingFinishedevFishingCancelevNewFloatObjectevNewFishingZoneObjectevFishingMiniGameevSteamAchievementCompletedevUpdatePuppetevChangeFlaggingFinishedevNewOutpostObjectevOutpostUpdateevOutpostClaimedevOutpostRewardevOverChargeEndevOverChargeStatusevPartyFinderFullUpdateevPartyFinderUpdateevPartyFinderApplicantsUpdateevPartyFinderEquipmentSnapshotevPartyFinderJoinRequestDeclinedevNewUnlockedPersonalSeasonRewardsevPersonalSeasonPointsGainedevEasyAntiCheatMessageToClientevMatchLootChestOpeningStartevMatchLootChestOpeningFinishedevMatchLootChestOpeningCancelevNotifyCrystalMatchRewardevCrystalRealmFeedbackevNewLocationMarkerevNewTutorialBlockerevNewTileSwitchevNewInformationProviderevNewDynamicGuildLogoevTutorialUpdateevTriggerHintBoxevRandomDungeonPositionInfoevNewLootChestevUpdateLootChestevLootChestOpenedevNewShrineevUpdateShrineevMutePlayerUpdateevShopTileUpdateevShopUpdateevEasyAntiCheatKickevUnlockVanityUnlockevAvatarUnlockedevCustomizationChangedevBaseVaultInfoevGuildVaultInfoevBankVaultInfoevRecoveryVaultPlayerInfoevRecoveryVaultGuildInfoevUpdateWardrobeevCastlePhaseChangedevGuildAccountLogEventevNewHideoutObjectevNewHideoutManagementevNewHideoutExitevInitHideoutAttackStartevInitHideoutAttackCancelevInitHideoutAttackFinishedevHideoutManagementUpdateevIpChangedevSmartClusterQueueUpdateInfoevSmartClusterQueueActiveInfoevSmartClusterQueueKickWarningevSmartClusterQueueInviteevReceivedGvgSeasonPointsevTerritoryBonusLevelUpdateevOpenWorldAttackScheduleStartevOpenWorldAttackScheduleFinishedevOpenWorldAttackScheduleCancelevOpenWorldAttackConquerStartevOpenWorldAttackConquerFinishedevOpenWorldAttackConquerCancelevOpenWorldAttackConquerStatusevOpenWorldAttackStartevOpenWorldAttackEndevNewRandomResourceBlockerevNewHomeObjectevHideoutObjectUpdateevUpdateInfamyevUnknown408evUnknown409evUnknown410evUnknown411evUnknown412evUnknown413evUnknown414evUnknown415evUnknown416evUnknown417evUnknown418evUnknown419evUnknown420evUnknown421evUnknown422"

var _EventType_index = [...]uint16{0, 8, 15, 29, 35, 45, 62, 76, 90, 110, 131, 157, 173, 181, 192, 210, 222, 238, 252, 263, 272, 282, 299, 315, 333, 354, 368, 386, 401, 419, 435, 451, 479, 511, 533, 550, 563, 587, 603, 624, 643, 663, 681, 703, 723, 741, 763, 782, 804, 829, 843, 858, 875, 887, 910, 934, 960, 984, 1009, 1036, 1058, 1081, 1106, 1125, 1139, 1152, 1161, 1174, 1185, 1196, 1207, 1222, 1242, 1255, 1267, 1289, 1309, 1325, 1348, 1357, 1373, 1400, 1427, 1454, 1486, 1515, 1553, 1585, 1604, 1613, 1634, 1655, 1680, 1699, 1712, 1732, 1748, 1772, 1792, 1805, 1824, 1844, 1858, 1875, 1895, 1918, 1933, 1951, 1969, 1981, 2007, 2029, 2048, 2063, 2076, 2095, 2103, 2119, 2139, 2160, 2188, 2216, 2231, 2249, 2265, 2276, 2288, 2306, 2331, 2345, 2373, 2398, 2417, 2433, 2455, 2471, 2498, 2526, 2538, 2558, 2579, 2600, 2625, 2654, 2682, 2712, 2736, 2757, 2771, 2781, 2795, 2812, 2823, 2837, 2843, 2856, 2880, 2903, 2934, 2959, 2990, 3021, 3044, 3062, 3081, 3100, 3121, 3146, 3159, 3184, 3209, 3226, 3239, 3250, 3262, 3276, 3295, 3310, 3342, 3364, 3381, 3399, 3421, 3449, 3482, 3516, 3536, 3552, 3564, 3587, 3600, 3617, 3636, 3653, 3673, 3687, 3696, 3708, 3721, 3737, 3759, 3768, 3780, 3800, 3819, 3838, 3860, 3886, 3909, 3936, 3961, 3982, 4000, 4007, 4024, 4037, 4053, 4072, 4091, 4108, 4128, 4159, 4178, 4198, 4225, 4252, 4279, 4297, 4318, 4331, 4348, 4367, 4390, 4406, 4423, 4442, 4465, 4486, 4501, 4520, 4545, 4565, 4594, 4615, 4643, 4668, 4694, 4709, 4732, 4752, 4773, 4797, 4822, 4843, 4857, 4878, 4896, 4917, 4939, 4963, 4988, 5013, 5041, 5067, 5102, 5136, 5152, 5169, 5198, 5218, 5234, 5252, 5267, 5287, 5300, 5323, 5343, 5366, 5381, 5396, 5412, 5435, 5453, 5477, 5499, 5524, 5551, 5580, 5596, 5615, 5636, 5656, 5681, 5703, 5714, 5727, 5742, 5757, 5770, 5789, 5804, 5826, 5846, 5872, 5886, 5900, 5926, 5947, 5971, 5991, 6007, 6025, 6038, 6059, 6081, 6105, 6130, 6150, 6174, 6202, 6225, 6242, 6265, 6295, 6309, 6322, 6336, 6353, 6368, 6384, 6406, 6423, 6450, 6464, 6488, 6506, 6521, 6537, 6552, 6567, 6585, 6608, 6627, 6656, 6686, 6718, 6752, 6780, 6810, 6838, 6869, 6898, 6924, 6946, 6965, 6985, 7000, 7024, 7045, 7061, 7077, 7104, 7118, 7135, 7152, 7163, 7177, 7195, 7211, 7223, 7242, 7262, 7278, 7300, 7315, 7331, 7346, 7371, 7395, 7411, 7431, 7453, 7471, 7493, 7509, 7533, 7558, 7585, 7610, 7621, 7650, 7679, 7709, 7734, 7759, 7786, 7816, 7849, 7880, 7909, 7941, 7971, 8001, 8023, 8043, 8069, 8084, 8105, 8119, 8131, 8143, 8155, 8167, 8179, 8191, 8203, 8215, 8227, 8239, 8251, 8263, 8275, 8287, 8299}

func (i EventType) String() string {
	if i >= EventType(len(_EventType_index)-1) {
		return "EventType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _EventType_name[_EventType_index[i]:_EventType_index[i+1]]
}
