package parser // Java20Parser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by Java20Parser.
type Java20ParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by Java20Parser#start_.
	VisitStart_(ctx *Start_Context) interface{}

	// Visit a parse tree produced by Java20Parser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by Java20Parser#typeIdentifier.
	VisitTypeIdentifier(ctx *TypeIdentifierContext) interface{}

	// Visit a parse tree produced by Java20Parser#unqualifiedMethodIdentifier.
	VisitUnqualifiedMethodIdentifier(ctx *UnqualifiedMethodIdentifierContext) interface{}

	// Visit a parse tree produced by Java20Parser#primitiveType.
	VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{}

	// Visit a parse tree produced by Java20Parser#numericType.
	VisitNumericType(ctx *NumericTypeContext) interface{}

	// Visit a parse tree produced by Java20Parser#integralType.
	VisitIntegralType(ctx *IntegralTypeContext) interface{}

	// Visit a parse tree produced by Java20Parser#floatingPointType.
	VisitFloatingPointType(ctx *FloatingPointTypeContext) interface{}

	// Visit a parse tree produced by Java20Parser#referenceType.
	VisitReferenceType(ctx *ReferenceTypeContext) interface{}

	// Visit a parse tree produced by Java20Parser#coit.
	VisitCoit(ctx *CoitContext) interface{}

	// Visit a parse tree produced by Java20Parser#classOrInterfaceType.
	VisitClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) interface{}

	// Visit a parse tree produced by Java20Parser#classType.
	VisitClassType(ctx *ClassTypeContext) interface{}

	// Visit a parse tree produced by Java20Parser#interfaceType.
	VisitInterfaceType(ctx *InterfaceTypeContext) interface{}

	// Visit a parse tree produced by Java20Parser#typeVariable.
	VisitTypeVariable(ctx *TypeVariableContext) interface{}

	// Visit a parse tree produced by Java20Parser#arrayType.
	VisitArrayType(ctx *ArrayTypeContext) interface{}

	// Visit a parse tree produced by Java20Parser#dims.
	VisitDims(ctx *DimsContext) interface{}

	// Visit a parse tree produced by Java20Parser#typeParameter.
	VisitTypeParameter(ctx *TypeParameterContext) interface{}

	// Visit a parse tree produced by Java20Parser#typeParameterModifier.
	VisitTypeParameterModifier(ctx *TypeParameterModifierContext) interface{}

	// Visit a parse tree produced by Java20Parser#typeBound.
	VisitTypeBound(ctx *TypeBoundContext) interface{}

	// Visit a parse tree produced by Java20Parser#additionalBound.
	VisitAdditionalBound(ctx *AdditionalBoundContext) interface{}

	// Visit a parse tree produced by Java20Parser#typeArguments.
	VisitTypeArguments(ctx *TypeArgumentsContext) interface{}

	// Visit a parse tree produced by Java20Parser#typeArgumentList.
	VisitTypeArgumentList(ctx *TypeArgumentListContext) interface{}

	// Visit a parse tree produced by Java20Parser#typeArgument.
	VisitTypeArgument(ctx *TypeArgumentContext) interface{}

	// Visit a parse tree produced by Java20Parser#wildcard.
	VisitWildcard(ctx *WildcardContext) interface{}

	// Visit a parse tree produced by Java20Parser#wildcardBounds.
	VisitWildcardBounds(ctx *WildcardBoundsContext) interface{}

	// Visit a parse tree produced by Java20Parser#moduleName.
	VisitModuleName(ctx *ModuleNameContext) interface{}

	// Visit a parse tree produced by Java20Parser#packageName.
	VisitPackageName(ctx *PackageNameContext) interface{}

	// Visit a parse tree produced by Java20Parser#typeName.
	VisitTypeName(ctx *TypeNameContext) interface{}

	// Visit a parse tree produced by Java20Parser#packageOrTypeName.
	VisitPackageOrTypeName(ctx *PackageOrTypeNameContext) interface{}

	// Visit a parse tree produced by Java20Parser#expressionName.
	VisitExpressionName(ctx *ExpressionNameContext) interface{}

	// Visit a parse tree produced by Java20Parser#methodName.
	VisitMethodName(ctx *MethodNameContext) interface{}

	// Visit a parse tree produced by Java20Parser#ambiguousName.
	VisitAmbiguousName(ctx *AmbiguousNameContext) interface{}

	// Visit a parse tree produced by Java20Parser#compilationUnit.
	VisitCompilationUnit(ctx *CompilationUnitContext) interface{}

	// Visit a parse tree produced by Java20Parser#ordinaryCompilationUnit.
	VisitOrdinaryCompilationUnit(ctx *OrdinaryCompilationUnitContext) interface{}

	// Visit a parse tree produced by Java20Parser#modularCompilationUnit.
	VisitModularCompilationUnit(ctx *ModularCompilationUnitContext) interface{}

	// Visit a parse tree produced by Java20Parser#packageDeclaration.
	VisitPackageDeclaration(ctx *PackageDeclarationContext) interface{}

	// Visit a parse tree produced by Java20Parser#packageModifier.
	VisitPackageModifier(ctx *PackageModifierContext) interface{}

	// Visit a parse tree produced by Java20Parser#importDeclaration.
	VisitImportDeclaration(ctx *ImportDeclarationContext) interface{}

	// Visit a parse tree produced by Java20Parser#singleTypeImportDeclaration.
	VisitSingleTypeImportDeclaration(ctx *SingleTypeImportDeclarationContext) interface{}

	// Visit a parse tree produced by Java20Parser#typeImportOnDemandDeclaration.
	VisitTypeImportOnDemandDeclaration(ctx *TypeImportOnDemandDeclarationContext) interface{}

	// Visit a parse tree produced by Java20Parser#singleStaticImportDeclaration.
	VisitSingleStaticImportDeclaration(ctx *SingleStaticImportDeclarationContext) interface{}

	// Visit a parse tree produced by Java20Parser#staticImportOnDemandDeclaration.
	VisitStaticImportOnDemandDeclaration(ctx *StaticImportOnDemandDeclarationContext) interface{}

	// Visit a parse tree produced by Java20Parser#topLevelClassOrInterfaceDeclaration.
	VisitTopLevelClassOrInterfaceDeclaration(ctx *TopLevelClassOrInterfaceDeclarationContext) interface{}

	// Visit a parse tree produced by Java20Parser#moduleDeclaration.
	VisitModuleDeclaration(ctx *ModuleDeclarationContext) interface{}

	// Visit a parse tree produced by Java20Parser#moduleDirective.
	VisitModuleDirective(ctx *ModuleDirectiveContext) interface{}

	// Visit a parse tree produced by Java20Parser#requiresModifier.
	VisitRequiresModifier(ctx *RequiresModifierContext) interface{}

	// Visit a parse tree produced by Java20Parser#classDeclaration.
	VisitClassDeclaration(ctx *ClassDeclarationContext) interface{}

	// Visit a parse tree produced by Java20Parser#normalClassDeclaration.
	VisitNormalClassDeclaration(ctx *NormalClassDeclarationContext) interface{}

	// Visit a parse tree produced by Java20Parser#classModifier.
	VisitClassModifier(ctx *ClassModifierContext) interface{}

	// Visit a parse tree produced by Java20Parser#typeParameters.
	VisitTypeParameters(ctx *TypeParametersContext) interface{}

	// Visit a parse tree produced by Java20Parser#typeParameterList.
	VisitTypeParameterList(ctx *TypeParameterListContext) interface{}

	// Visit a parse tree produced by Java20Parser#classExtends.
	VisitClassExtends(ctx *ClassExtendsContext) interface{}

	// Visit a parse tree produced by Java20Parser#classImplements.
	VisitClassImplements(ctx *ClassImplementsContext) interface{}

	// Visit a parse tree produced by Java20Parser#interfaceTypeList.
	VisitInterfaceTypeList(ctx *InterfaceTypeListContext) interface{}

	// Visit a parse tree produced by Java20Parser#classPermits.
	VisitClassPermits(ctx *ClassPermitsContext) interface{}

	// Visit a parse tree produced by Java20Parser#classBody.
	VisitClassBody(ctx *ClassBodyContext) interface{}

	// Visit a parse tree produced by Java20Parser#classBodyDeclaration.
	VisitClassBodyDeclaration(ctx *ClassBodyDeclarationContext) interface{}

	// Visit a parse tree produced by Java20Parser#classMemberDeclaration.
	VisitClassMemberDeclaration(ctx *ClassMemberDeclarationContext) interface{}

	// Visit a parse tree produced by Java20Parser#fieldDeclaration.
	VisitFieldDeclaration(ctx *FieldDeclarationContext) interface{}

	// Visit a parse tree produced by Java20Parser#fieldModifier.
	VisitFieldModifier(ctx *FieldModifierContext) interface{}

	// Visit a parse tree produced by Java20Parser#variableDeclaratorList.
	VisitVariableDeclaratorList(ctx *VariableDeclaratorListContext) interface{}

	// Visit a parse tree produced by Java20Parser#variableDeclarator.
	VisitVariableDeclarator(ctx *VariableDeclaratorContext) interface{}

	// Visit a parse tree produced by Java20Parser#variableDeclaratorId.
	VisitVariableDeclaratorId(ctx *VariableDeclaratorIdContext) interface{}

	// Visit a parse tree produced by Java20Parser#variableInitializer.
	VisitVariableInitializer(ctx *VariableInitializerContext) interface{}

	// Visit a parse tree produced by Java20Parser#unannType.
	VisitUnannType(ctx *UnannTypeContext) interface{}

	// Visit a parse tree produced by Java20Parser#unannPrimitiveType.
	VisitUnannPrimitiveType(ctx *UnannPrimitiveTypeContext) interface{}

	// Visit a parse tree produced by Java20Parser#unannReferenceType.
	VisitUnannReferenceType(ctx *UnannReferenceTypeContext) interface{}

	// Visit a parse tree produced by Java20Parser#unannClassOrInterfaceType.
	VisitUnannClassOrInterfaceType(ctx *UnannClassOrInterfaceTypeContext) interface{}

	// Visit a parse tree produced by Java20Parser#uCOIT.
	VisitUCOIT(ctx *UCOITContext) interface{}

	// Visit a parse tree produced by Java20Parser#unannClassType.
	VisitUnannClassType(ctx *UnannClassTypeContext) interface{}

	// Visit a parse tree produced by Java20Parser#unannInterfaceType.
	VisitUnannInterfaceType(ctx *UnannInterfaceTypeContext) interface{}

	// Visit a parse tree produced by Java20Parser#unannTypeVariable.
	VisitUnannTypeVariable(ctx *UnannTypeVariableContext) interface{}

	// Visit a parse tree produced by Java20Parser#unannArrayType.
	VisitUnannArrayType(ctx *UnannArrayTypeContext) interface{}

	// Visit a parse tree produced by Java20Parser#methodDeclaration.
	VisitMethodDeclaration(ctx *MethodDeclarationContext) interface{}

	// Visit a parse tree produced by Java20Parser#methodModifier.
	VisitMethodModifier(ctx *MethodModifierContext) interface{}

	// Visit a parse tree produced by Java20Parser#methodHeader.
	VisitMethodHeader(ctx *MethodHeaderContext) interface{}

	// Visit a parse tree produced by Java20Parser#result.
	VisitResult(ctx *ResultContext) interface{}

	// Visit a parse tree produced by Java20Parser#methodDeclarator.
	VisitMethodDeclarator(ctx *MethodDeclaratorContext) interface{}

	// Visit a parse tree produced by Java20Parser#receiverParameter.
	VisitReceiverParameter(ctx *ReceiverParameterContext) interface{}

	// Visit a parse tree produced by Java20Parser#formalParameterList.
	VisitFormalParameterList(ctx *FormalParameterListContext) interface{}

	// Visit a parse tree produced by Java20Parser#formalParameter.
	VisitFormalParameter(ctx *FormalParameterContext) interface{}

	// Visit a parse tree produced by Java20Parser#variableArityParameter.
	VisitVariableArityParameter(ctx *VariableArityParameterContext) interface{}

	// Visit a parse tree produced by Java20Parser#variableModifier.
	VisitVariableModifier(ctx *VariableModifierContext) interface{}

	// Visit a parse tree produced by Java20Parser#throwsT.
	VisitThrowsT(ctx *ThrowsTContext) interface{}

	// Visit a parse tree produced by Java20Parser#exceptionTypeList.
	VisitExceptionTypeList(ctx *ExceptionTypeListContext) interface{}

	// Visit a parse tree produced by Java20Parser#exceptionType.
	VisitExceptionType(ctx *ExceptionTypeContext) interface{}

	// Visit a parse tree produced by Java20Parser#methodBody.
	VisitMethodBody(ctx *MethodBodyContext) interface{}

	// Visit a parse tree produced by Java20Parser#instanceInitializer.
	VisitInstanceInitializer(ctx *InstanceInitializerContext) interface{}

	// Visit a parse tree produced by Java20Parser#staticInitializer.
	VisitStaticInitializer(ctx *StaticInitializerContext) interface{}

	// Visit a parse tree produced by Java20Parser#constructorDeclaration.
	VisitConstructorDeclaration(ctx *ConstructorDeclarationContext) interface{}

	// Visit a parse tree produced by Java20Parser#constructorModifier.
	VisitConstructorModifier(ctx *ConstructorModifierContext) interface{}

	// Visit a parse tree produced by Java20Parser#constructorDeclarator.
	VisitConstructorDeclarator(ctx *ConstructorDeclaratorContext) interface{}

	// Visit a parse tree produced by Java20Parser#simpleTypeName.
	VisitSimpleTypeName(ctx *SimpleTypeNameContext) interface{}

	// Visit a parse tree produced by Java20Parser#constructorBody.
	VisitConstructorBody(ctx *ConstructorBodyContext) interface{}

	// Visit a parse tree produced by Java20Parser#explicitConstructorInvocation.
	VisitExplicitConstructorInvocation(ctx *ExplicitConstructorInvocationContext) interface{}

	// Visit a parse tree produced by Java20Parser#enumDeclaration.
	VisitEnumDeclaration(ctx *EnumDeclarationContext) interface{}

	// Visit a parse tree produced by Java20Parser#enumBody.
	VisitEnumBody(ctx *EnumBodyContext) interface{}

	// Visit a parse tree produced by Java20Parser#enumConstantList.
	VisitEnumConstantList(ctx *EnumConstantListContext) interface{}

	// Visit a parse tree produced by Java20Parser#enumConstant.
	VisitEnumConstant(ctx *EnumConstantContext) interface{}

	// Visit a parse tree produced by Java20Parser#enumConstantModifier.
	VisitEnumConstantModifier(ctx *EnumConstantModifierContext) interface{}

	// Visit a parse tree produced by Java20Parser#enumBodyDeclarations.
	VisitEnumBodyDeclarations(ctx *EnumBodyDeclarationsContext) interface{}

	// Visit a parse tree produced by Java20Parser#recordDeclaration.
	VisitRecordDeclaration(ctx *RecordDeclarationContext) interface{}

	// Visit a parse tree produced by Java20Parser#recordHeader.
	VisitRecordHeader(ctx *RecordHeaderContext) interface{}

	// Visit a parse tree produced by Java20Parser#recordComponentList.
	VisitRecordComponentList(ctx *RecordComponentListContext) interface{}

	// Visit a parse tree produced by Java20Parser#recordComponent.
	VisitRecordComponent(ctx *RecordComponentContext) interface{}

	// Visit a parse tree produced by Java20Parser#variableArityRecordComponent.
	VisitVariableArityRecordComponent(ctx *VariableArityRecordComponentContext) interface{}

	// Visit a parse tree produced by Java20Parser#recordComponentModifier.
	VisitRecordComponentModifier(ctx *RecordComponentModifierContext) interface{}

	// Visit a parse tree produced by Java20Parser#recordBody.
	VisitRecordBody(ctx *RecordBodyContext) interface{}

	// Visit a parse tree produced by Java20Parser#recordBodyDeclaration.
	VisitRecordBodyDeclaration(ctx *RecordBodyDeclarationContext) interface{}

	// Visit a parse tree produced by Java20Parser#compactConstructorDeclaration.
	VisitCompactConstructorDeclaration(ctx *CompactConstructorDeclarationContext) interface{}

	// Visit a parse tree produced by Java20Parser#interfaceDeclaration.
	VisitInterfaceDeclaration(ctx *InterfaceDeclarationContext) interface{}

	// Visit a parse tree produced by Java20Parser#normalInterfaceDeclaration.
	VisitNormalInterfaceDeclaration(ctx *NormalInterfaceDeclarationContext) interface{}

	// Visit a parse tree produced by Java20Parser#interfaceModifier.
	VisitInterfaceModifier(ctx *InterfaceModifierContext) interface{}

	// Visit a parse tree produced by Java20Parser#interfaceExtends.
	VisitInterfaceExtends(ctx *InterfaceExtendsContext) interface{}

	// Visit a parse tree produced by Java20Parser#interfacePermits.
	VisitInterfacePermits(ctx *InterfacePermitsContext) interface{}

	// Visit a parse tree produced by Java20Parser#interfaceBody.
	VisitInterfaceBody(ctx *InterfaceBodyContext) interface{}

	// Visit a parse tree produced by Java20Parser#interfaceMemberDeclaration.
	VisitInterfaceMemberDeclaration(ctx *InterfaceMemberDeclarationContext) interface{}

	// Visit a parse tree produced by Java20Parser#constantDeclaration.
	VisitConstantDeclaration(ctx *ConstantDeclarationContext) interface{}

	// Visit a parse tree produced by Java20Parser#constantModifier.
	VisitConstantModifier(ctx *ConstantModifierContext) interface{}

	// Visit a parse tree produced by Java20Parser#interfaceMethodDeclaration.
	VisitInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) interface{}

	// Visit a parse tree produced by Java20Parser#interfaceMethodModifier.
	VisitInterfaceMethodModifier(ctx *InterfaceMethodModifierContext) interface{}

	// Visit a parse tree produced by Java20Parser#annotationInterfaceDeclaration.
	VisitAnnotationInterfaceDeclaration(ctx *AnnotationInterfaceDeclarationContext) interface{}

	// Visit a parse tree produced by Java20Parser#annotationInterfaceBody.
	VisitAnnotationInterfaceBody(ctx *AnnotationInterfaceBodyContext) interface{}

	// Visit a parse tree produced by Java20Parser#annotationInterfaceMemberDeclaration.
	VisitAnnotationInterfaceMemberDeclaration(ctx *AnnotationInterfaceMemberDeclarationContext) interface{}

	// Visit a parse tree produced by Java20Parser#annotationInterfaceElementDeclaration.
	VisitAnnotationInterfaceElementDeclaration(ctx *AnnotationInterfaceElementDeclarationContext) interface{}

	// Visit a parse tree produced by Java20Parser#annotationInterfaceElementModifier.
	VisitAnnotationInterfaceElementModifier(ctx *AnnotationInterfaceElementModifierContext) interface{}

	// Visit a parse tree produced by Java20Parser#defaultValue.
	VisitDefaultValue(ctx *DefaultValueContext) interface{}

	// Visit a parse tree produced by Java20Parser#annotation.
	VisitAnnotation(ctx *AnnotationContext) interface{}

	// Visit a parse tree produced by Java20Parser#normalAnnotation.
	VisitNormalAnnotation(ctx *NormalAnnotationContext) interface{}

	// Visit a parse tree produced by Java20Parser#elementValuePairList.
	VisitElementValuePairList(ctx *ElementValuePairListContext) interface{}

	// Visit a parse tree produced by Java20Parser#elementValuePair.
	VisitElementValuePair(ctx *ElementValuePairContext) interface{}

	// Visit a parse tree produced by Java20Parser#elementValue.
	VisitElementValue(ctx *ElementValueContext) interface{}

	// Visit a parse tree produced by Java20Parser#elementValueArrayInitializer.
	VisitElementValueArrayInitializer(ctx *ElementValueArrayInitializerContext) interface{}

	// Visit a parse tree produced by Java20Parser#elementValueList.
	VisitElementValueList(ctx *ElementValueListContext) interface{}

	// Visit a parse tree produced by Java20Parser#markerAnnotation.
	VisitMarkerAnnotation(ctx *MarkerAnnotationContext) interface{}

	// Visit a parse tree produced by Java20Parser#singleElementAnnotation.
	VisitSingleElementAnnotation(ctx *SingleElementAnnotationContext) interface{}

	// Visit a parse tree produced by Java20Parser#arrayInitializer.
	VisitArrayInitializer(ctx *ArrayInitializerContext) interface{}

	// Visit a parse tree produced by Java20Parser#variableInitializerList.
	VisitVariableInitializerList(ctx *VariableInitializerListContext) interface{}

	// Visit a parse tree produced by Java20Parser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by Java20Parser#blockStatements.
	VisitBlockStatements(ctx *BlockStatementsContext) interface{}

	// Visit a parse tree produced by Java20Parser#blockStatement.
	VisitBlockStatement(ctx *BlockStatementContext) interface{}

	// Visit a parse tree produced by Java20Parser#localClassOrInterfaceDeclaration.
	VisitLocalClassOrInterfaceDeclaration(ctx *LocalClassOrInterfaceDeclarationContext) interface{}

	// Visit a parse tree produced by Java20Parser#localVariableDeclaration.
	VisitLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) interface{}

	// Visit a parse tree produced by Java20Parser#localVariableType.
	VisitLocalVariableType(ctx *LocalVariableTypeContext) interface{}

	// Visit a parse tree produced by Java20Parser#localVariableDeclarationStatement.
	VisitLocalVariableDeclarationStatement(ctx *LocalVariableDeclarationStatementContext) interface{}

	// Visit a parse tree produced by Java20Parser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by Java20Parser#statementNoShortIf.
	VisitStatementNoShortIf(ctx *StatementNoShortIfContext) interface{}

	// Visit a parse tree produced by Java20Parser#statementWithoutTrailingSubstatement.
	VisitStatementWithoutTrailingSubstatement(ctx *StatementWithoutTrailingSubstatementContext) interface{}

	// Visit a parse tree produced by Java20Parser#emptyStatement_.
	VisitEmptyStatement_(ctx *EmptyStatement_Context) interface{}

	// Visit a parse tree produced by Java20Parser#labeledStatement.
	VisitLabeledStatement(ctx *LabeledStatementContext) interface{}

	// Visit a parse tree produced by Java20Parser#labeledStatementNoShortIf.
	VisitLabeledStatementNoShortIf(ctx *LabeledStatementNoShortIfContext) interface{}

	// Visit a parse tree produced by Java20Parser#expressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by Java20Parser#statementExpression.
	VisitStatementExpression(ctx *StatementExpressionContext) interface{}

	// Visit a parse tree produced by Java20Parser#ifThenStatement.
	VisitIfThenStatement(ctx *IfThenStatementContext) interface{}

	// Visit a parse tree produced by Java20Parser#ifThenElseStatement.
	VisitIfThenElseStatement(ctx *IfThenElseStatementContext) interface{}

	// Visit a parse tree produced by Java20Parser#ifThenElseStatementNoShortIf.
	VisitIfThenElseStatementNoShortIf(ctx *IfThenElseStatementNoShortIfContext) interface{}

	// Visit a parse tree produced by Java20Parser#assertStatement.
	VisitAssertStatement(ctx *AssertStatementContext) interface{}

	// Visit a parse tree produced by Java20Parser#switchStatement.
	VisitSwitchStatement(ctx *SwitchStatementContext) interface{}

	// Visit a parse tree produced by Java20Parser#switchBlock.
	VisitSwitchBlock(ctx *SwitchBlockContext) interface{}

	// Visit a parse tree produced by Java20Parser#switchRule.
	VisitSwitchRule(ctx *SwitchRuleContext) interface{}

	// Visit a parse tree produced by Java20Parser#switchBlockStatementGroup.
	VisitSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) interface{}

	// Visit a parse tree produced by Java20Parser#switchLabel.
	VisitSwitchLabel(ctx *SwitchLabelContext) interface{}

	// Visit a parse tree produced by Java20Parser#caseConstant.
	VisitCaseConstant(ctx *CaseConstantContext) interface{}

	// Visit a parse tree produced by Java20Parser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by Java20Parser#whileStatementNoShortIf.
	VisitWhileStatementNoShortIf(ctx *WhileStatementNoShortIfContext) interface{}

	// Visit a parse tree produced by Java20Parser#doStatement.
	VisitDoStatement(ctx *DoStatementContext) interface{}

	// Visit a parse tree produced by Java20Parser#forStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by Java20Parser#forStatementNoShortIf.
	VisitForStatementNoShortIf(ctx *ForStatementNoShortIfContext) interface{}

	// Visit a parse tree produced by Java20Parser#basicForStatement.
	VisitBasicForStatement(ctx *BasicForStatementContext) interface{}

	// Visit a parse tree produced by Java20Parser#basicForStatementNoShortIf.
	VisitBasicForStatementNoShortIf(ctx *BasicForStatementNoShortIfContext) interface{}

	// Visit a parse tree produced by Java20Parser#forInit.
	VisitForInit(ctx *ForInitContext) interface{}

	// Visit a parse tree produced by Java20Parser#forUpdate.
	VisitForUpdate(ctx *ForUpdateContext) interface{}

	// Visit a parse tree produced by Java20Parser#statementExpressionList.
	VisitStatementExpressionList(ctx *StatementExpressionListContext) interface{}

	// Visit a parse tree produced by Java20Parser#enhancedForStatement.
	VisitEnhancedForStatement(ctx *EnhancedForStatementContext) interface{}

	// Visit a parse tree produced by Java20Parser#enhancedForStatementNoShortIf.
	VisitEnhancedForStatementNoShortIf(ctx *EnhancedForStatementNoShortIfContext) interface{}

	// Visit a parse tree produced by Java20Parser#breakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by Java20Parser#continueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by Java20Parser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by Java20Parser#throwStatement.
	VisitThrowStatement(ctx *ThrowStatementContext) interface{}

	// Visit a parse tree produced by Java20Parser#synchronizedStatement.
	VisitSynchronizedStatement(ctx *SynchronizedStatementContext) interface{}

	// Visit a parse tree produced by Java20Parser#TryStatement1.
	VisitTryStatement1(ctx *TryStatement1Context) interface{}

	// Visit a parse tree produced by Java20Parser#TryStatement2.
	VisitTryStatement2(ctx *TryStatement2Context) interface{}

	// Visit a parse tree produced by Java20Parser#TryStatement3.
	VisitTryStatement3(ctx *TryStatement3Context) interface{}

	// Visit a parse tree produced by Java20Parser#TryStatement4.
	VisitTryStatement4(ctx *TryStatement4Context) interface{}

	// Visit a parse tree produced by Java20Parser#catches.
	VisitCatches(ctx *CatchesContext) interface{}

	// Visit a parse tree produced by Java20Parser#catchClause.
	VisitCatchClause(ctx *CatchClauseContext) interface{}

	// Visit a parse tree produced by Java20Parser#catchFormalParameter.
	VisitCatchFormalParameter(ctx *CatchFormalParameterContext) interface{}

	// Visit a parse tree produced by Java20Parser#catchType.
	VisitCatchType(ctx *CatchTypeContext) interface{}

	// Visit a parse tree produced by Java20Parser#finallyBlock.
	VisitFinallyBlock(ctx *FinallyBlockContext) interface{}

	// Visit a parse tree produced by Java20Parser#tryWithResourcesStatement.
	VisitTryWithResourcesStatement(ctx *TryWithResourcesStatementContext) interface{}

	// Visit a parse tree produced by Java20Parser#resourceSpecification.
	VisitResourceSpecification(ctx *ResourceSpecificationContext) interface{}

	// Visit a parse tree produced by Java20Parser#resourceList.
	VisitResourceList(ctx *ResourceListContext) interface{}

	// Visit a parse tree produced by Java20Parser#resource.
	VisitResource(ctx *ResourceContext) interface{}

	// Visit a parse tree produced by Java20Parser#variableAccess.
	VisitVariableAccess(ctx *VariableAccessContext) interface{}

	// Visit a parse tree produced by Java20Parser#yieldStatement.
	VisitYieldStatement(ctx *YieldStatementContext) interface{}

	// Visit a parse tree produced by Java20Parser#pattern.
	VisitPattern(ctx *PatternContext) interface{}

	// Visit a parse tree produced by Java20Parser#typePattern.
	VisitTypePattern(ctx *TypePatternContext) interface{}

	// Visit a parse tree produced by Java20Parser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by Java20Parser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by Java20Parser#primaryNoNewArray.
	VisitPrimaryNoNewArray(ctx *PrimaryNoNewArrayContext) interface{}

	// Visit a parse tree produced by Java20Parser#pNNA.
	VisitPNNA(ctx *PNNAContext) interface{}

	// Visit a parse tree produced by Java20Parser#classLiteral.
	VisitClassLiteral(ctx *ClassLiteralContext) interface{}

	// Visit a parse tree produced by Java20Parser#classInstanceCreationExpression.
	VisitClassInstanceCreationExpression(ctx *ClassInstanceCreationExpressionContext) interface{}

	// Visit a parse tree produced by Java20Parser#unqualifiedClassInstanceCreationExpression.
	VisitUnqualifiedClassInstanceCreationExpression(ctx *UnqualifiedClassInstanceCreationExpressionContext) interface{}

	// Visit a parse tree produced by Java20Parser#classOrInterfaceTypeToInstantiate.
	VisitClassOrInterfaceTypeToInstantiate(ctx *ClassOrInterfaceTypeToInstantiateContext) interface{}

	// Visit a parse tree produced by Java20Parser#typeArgumentsOrDiamond.
	VisitTypeArgumentsOrDiamond(ctx *TypeArgumentsOrDiamondContext) interface{}

	// Visit a parse tree produced by Java20Parser#arrayCreationExpression.
	VisitArrayCreationExpression(ctx *ArrayCreationExpressionContext) interface{}

	// Visit a parse tree produced by Java20Parser#arrayCreationExpressionWithoutInitializer.
	VisitArrayCreationExpressionWithoutInitializer(ctx *ArrayCreationExpressionWithoutInitializerContext) interface{}

	// Visit a parse tree produced by Java20Parser#arrayCreationExpressionWithInitializer.
	VisitArrayCreationExpressionWithInitializer(ctx *ArrayCreationExpressionWithInitializerContext) interface{}

	// Visit a parse tree produced by Java20Parser#dimExprs.
	VisitDimExprs(ctx *DimExprsContext) interface{}

	// Visit a parse tree produced by Java20Parser#dimExpr.
	VisitDimExpr(ctx *DimExprContext) interface{}

	// Visit a parse tree produced by Java20Parser#arrayAccess.
	VisitArrayAccess(ctx *ArrayAccessContext) interface{}

	// Visit a parse tree produced by Java20Parser#fieldAccess.
	VisitFieldAccess(ctx *FieldAccessContext) interface{}

	// Visit a parse tree produced by Java20Parser#methodInvocation.
	VisitMethodInvocation(ctx *MethodInvocationContext) interface{}

	// Visit a parse tree produced by Java20Parser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by Java20Parser#methodReference.
	VisitMethodReference(ctx *MethodReferenceContext) interface{}

	// Visit a parse tree produced by Java20Parser#postfixExpression.
	VisitPostfixExpression(ctx *PostfixExpressionContext) interface{}

	// Visit a parse tree produced by Java20Parser#pfE.
	VisitPfE(ctx *PfEContext) interface{}

	// Visit a parse tree produced by Java20Parser#postIncrementExpression.
	VisitPostIncrementExpression(ctx *PostIncrementExpressionContext) interface{}

	// Visit a parse tree produced by Java20Parser#postDecrementExpression.
	VisitPostDecrementExpression(ctx *PostDecrementExpressionContext) interface{}

	// Visit a parse tree produced by Java20Parser#unaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by Java20Parser#preIncrementExpression.
	VisitPreIncrementExpression(ctx *PreIncrementExpressionContext) interface{}

	// Visit a parse tree produced by Java20Parser#preDecrementExpression.
	VisitPreDecrementExpression(ctx *PreDecrementExpressionContext) interface{}

	// Visit a parse tree produced by Java20Parser#unaryExpressionNotPlusMinus.
	VisitUnaryExpressionNotPlusMinus(ctx *UnaryExpressionNotPlusMinusContext) interface{}

	// Visit a parse tree produced by Java20Parser#castExpression.
	VisitCastExpression(ctx *CastExpressionContext) interface{}

	// Visit a parse tree produced by Java20Parser#multiplicativeExpression.
	VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{}

	// Visit a parse tree produced by Java20Parser#additiveExpression.
	VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{}

	// Visit a parse tree produced by Java20Parser#shiftExpression.
	VisitShiftExpression(ctx *ShiftExpressionContext) interface{}

	// Visit a parse tree produced by Java20Parser#relationalExpression.
	VisitRelationalExpression(ctx *RelationalExpressionContext) interface{}

	// Visit a parse tree produced by Java20Parser#equalityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by Java20Parser#andExpression.
	VisitAndExpression(ctx *AndExpressionContext) interface{}

	// Visit a parse tree produced by Java20Parser#exclusiveOrExpression.
	VisitExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) interface{}

	// Visit a parse tree produced by Java20Parser#inclusiveOrExpression.
	VisitInclusiveOrExpression(ctx *InclusiveOrExpressionContext) interface{}

	// Visit a parse tree produced by Java20Parser#conditionalAndExpression.
	VisitConditionalAndExpression(ctx *ConditionalAndExpressionContext) interface{}

	// Visit a parse tree produced by Java20Parser#conditionalOrExpression.
	VisitConditionalOrExpression(ctx *ConditionalOrExpressionContext) interface{}

	// Visit a parse tree produced by Java20Parser#conditionalExpression.
	VisitConditionalExpression(ctx *ConditionalExpressionContext) interface{}

	// Visit a parse tree produced by Java20Parser#assignmentExpression.
	VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{}

	// Visit a parse tree produced by Java20Parser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by Java20Parser#leftHandSide.
	VisitLeftHandSide(ctx *LeftHandSideContext) interface{}

	// Visit a parse tree produced by Java20Parser#assignmentOperator.
	VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{}

	// Visit a parse tree produced by Java20Parser#lambdaExpression.
	VisitLambdaExpression(ctx *LambdaExpressionContext) interface{}

	// Visit a parse tree produced by Java20Parser#lambdaParameters.
	VisitLambdaParameters(ctx *LambdaParametersContext) interface{}

	// Visit a parse tree produced by Java20Parser#lambdaParameterList.
	VisitLambdaParameterList(ctx *LambdaParameterListContext) interface{}

	// Visit a parse tree produced by Java20Parser#lambdaParameter.
	VisitLambdaParameter(ctx *LambdaParameterContext) interface{}

	// Visit a parse tree produced by Java20Parser#lambdaParameterType.
	VisitLambdaParameterType(ctx *LambdaParameterTypeContext) interface{}

	// Visit a parse tree produced by Java20Parser#lambdaBody.
	VisitLambdaBody(ctx *LambdaBodyContext) interface{}

	// Visit a parse tree produced by Java20Parser#switchExpression.
	VisitSwitchExpression(ctx *SwitchExpressionContext) interface{}

	// Visit a parse tree produced by Java20Parser#constantExpression.
	VisitConstantExpression(ctx *ConstantExpressionContext) interface{}
}
